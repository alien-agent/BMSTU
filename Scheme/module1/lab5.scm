(load "unit-test.rkt")
(define (end? program count)
  (if (equal? (vector-ref program count) 'end)
      (+ count 1)
      (end? program (+ count 1))))

(define (skip-if program count num-ifs)
  (cond
    ((and (equal? (vector-ref program count) 'endif)
          (equal? num-ifs 1))
     (+ count 1))
    ((equal? (vector-ref program count) 'endif)
     (skip-if program (+ count 1) (- num-ifs 1)))
    ((equal? (vector-ref program count) 'if)
     (skip-if program (+ count 1) (+ num-ifs 1)))
    (else
     (skip-if program (+ count 1) num-ifs))))

(define (interpret program given-stack)
  (define (main wcount stack return dict)
    (if (< wcount (vector-length program))
        (let* ((word (vector-ref program wcount))
               (func (assoc word dict)))
          (cond
            ((number? word)
             (main (+ wcount 1) (cons word stack) return dict))
            (func (main (cadr func) stack (cons (+ wcount 1) return) dict))
            ((equal? word 'define)
             (main (end? program wcount)
                   stack
                   return
                   (cons (list (vector-ref program (+ wcount 1)) (+ wcount 2)) dict)))
            ((or (equal? word 'end) 
                 (equal? word 'exit))             
             (main (car return) stack (cdr return) dict))            
            ((equal? word 'if)
             (main (if (equal? (car stack) 0)
                       (skip-if program (+ wcount 1) 1)
                       (+ wcount 1))
                   (cdr stack)
                   return
                   dict))            
            ((equal? word 'endif)
             (main (+ wcount 1) stack return dict))            
            ((equal? word '+)
             (main (+ wcount 1) (cons (+ (cadr stack) (car stack)) (cddr stack)) return dict))            
            ((equal? word '-)
             (main (+ wcount 1) (cons (- (cadr stack) (car stack)) (cddr stack)) return dict))            
            ((equal? word '*)
             (main (+ wcount 1) (cons (* (cadr stack) (car stack)) (cddr stack)) return dict))            
            ((equal? word '/)
             (main (+ wcount 1) (cons (quotient (cadr stack) (car stack)) (cddr stack)) return dict))            
            ((equal? word 'mod)
             (main (+ wcount 1) (cons (remainder (cadr stack) (car stack)) (cddr stack)) return dict))            
            ((equal? word 'neg)
             (main (+ wcount 1) (cons (- (car stack)) (cdr stack)) return dict))            
            ((equal? word '=)
             (main (+ wcount 1) (cons (if (equal? (cadr stack) (car stack)) -1 0) (cddr stack)) return dict))            
            ((equal? word '>)
             (main (+ wcount 1) (cons (if (> (cadr stack) (car stack)) -1 0) (cddr stack)) return dict))            
            ((equal? word '<)
             (main (+ wcount 1) (cons (if (< (cadr stack) (car stack)) -1 0) (cddr stack)) return dict))            
            ((equal? word 'not)
             (main (+ wcount 1) (cons (if (equal? (car stack) 0) -1 0) (cdr stack)) return dict))            
            ((equal? word 'and)
             (main (+ wcount 1) (cons (if (or (equal? (cadr stack) 0) (equal? (car stack) 0)) 0 -1) (cddr stack)) return dict))            
            ((equal? word 'or)
             (main (+ wcount 1) (cons (if (and (equal? (cadr stack) 0) (equal? (car stack) 0)) 0 -1) (cddr stack)) return dict))            
            ((equal? word 'drop) 
             (main (+ wcount 1) (cdr stack) return dict))            
            ((equal? word 'swap) 
             (main (+ wcount 1) (cons (cadr stack) (cons (car stack) (cddr stack))) return dict))            
            ((equal? word 'dup)
             (main (+ wcount 1) (cons (car stack) stack) return dict))            
            ((equal? word 'over)
             (main (+ wcount 1) (cons (cadr stack) stack) return dict))            
            ((equal? word 'rot)
             (main (+ wcount 1) (cons (caddr stack) (cons (cadr stack) (cons (car stack) (cdddr stack)))) return dict))            
            ((equal? word 'depth)
             (main (+ wcount 1) (cons (length stack) stack) return dict))))
        stack))
  (main 0 given-stack '() '()))

(define the-tests
  (list (test (interpret #(define abs 
                            dup 0 < 
                            if neg endif 
                            end 
                            9 abs 
                            -9 abs) (quote ())) '(9 9))
        (test (interpret #(define =0? dup 0 = end
                            define <0? dup 0 < end
                            define signum
                            =0? if exit endif
                            <0? if drop -1 exit endif
                            drop
                            1
                            end
                            0 signum
                            -5 signum
                            10 signum) (quote ())) '(1 -1 0))
        (test (interpret #(   define -- 1 - end
                               define =0? dup 0 = end
                               define =1? dup 1 = end
                               define factorial
                               =0? if drop 1 exit endif
                               =1? if drop 1 exit endif
                               dup --
                               factorial
                               *
                               end
                               0 factorial
                               1 factorial
                               2 factorial
                               3 factorial
                               4 factorial     ) (quote ())) '(24 6 2 1 1))
        (test (interpret #(   define =0? dup 0 = end
                               define =1? dup 1 = end
                               define -- 1 - end
                               define fib
                               =0? if drop 0 exit endif
                               =1? if drop 1 exit endif
                               -- dup
                               -- fib
                               swap fib
                               +
                               end
                               define make-fib
                               dup 0 < if drop exit endif
                               dup fib
                               swap --
                               make-fib
                               end
                               10 make-fib     ) (quote ())) '(0 1 1 2 3 5 8 13 21 34 55))
        (test (interpret #(   define =0? dup 0 = end
                               define gcd
                               =0? if drop exit endif
                               swap over mod
                               gcd
                               end
                               90 99 gcd
                               234 8100 gcd    ) '()) '(18 9))))

(run-tests the-tests)