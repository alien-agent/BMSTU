; Task 1
(define stop #f)

(define-syntax use-assertions
  (syntax-rules ()
    ((_) (call-with-current-continuation
          (lambda (expression)
            (begin
              (set! stop expression)))))))

(define-syntax assert
  (syntax-rules ()
    ((_ expr) (if (not expr)
                  (begin (display "FAILED: ")
                         (display 'expr)
                         (stop))))))

(use-assertions)

(define (1/x x)
  (assert (not (zero? x)))
  (/ 1 x))

(map 1/x '(1 2 3 4 5))
(map 1/x '(-2 -1 0 1 2))

; Task 2
(define (save-data data addr)
  (call-with-output-file addr
    (lambda (port)
      (write data port))))

(define (load-data addr)
  (call-with-input-file addr
    (lambda (port)
      (read port))))

(define (string-counter addr)
  (call-with-input-file addr
    (lambda (port)
      (define (helper counter)
        (if (eof-object? (peek-char port)) ; достигли конца файла -> возвращаем значение
            counter
            (if (equal? newline (read-char port)) ; считываем следующий символ
                (begin (read-char port) (helper (+ counter 1)))
                (begin (read-char port) (helper (+ counter 1))))))
      (helper 0))))
; Task 3

(define memotrib
  (let ((memo '()))
    (lambda (n)
      (if (assoc n memo)
          (cadr (assoc n memo))
          (let ((res (if (<= n 1)
                         0
                         (if (= n 2)
                             1
                             (+ (memotrib (- n 3)) (memotrib (- n 2)) (memotrib (- n 1)))))))
            (set! memo (cons (list n res) memo))
            res)))))

; Task 4
(define-syntax my-if
  (syntax-rules ()
    ((_ cond? exp1 exp2) (let ((x (delay exp1))
                               (y (delay exp2)))
                           (force (or (and cond? x) y))))))

; Task 5
(define-syntax my-let
  (syntax-rules ()
    ((_ ((id val-expr) ...) body ...)
     ((lambda (id ...) body ...) val-expr ...))))

(define-syntax my-let*
  (syntax-rules ()
    ((_ ((id val-expr)) body ...)
     (my-let ((id val-expr)) body ...))
    ((_ ((id val-expr) (id1 val-expr1) ...) body ...)
     (my-let ((id val-expr))
             (my-let* ((id1 val-expr1) ...) body ...)))))

; Task 6
(define-syntax when
  (syntax-rules()
    ((when cond? expr ...)
     (if cond? (begin expr ...)))))

;(define x 3)
;(when   (> x 0) (display x) (display " > 0 ") (newline))

(define-syntax unless
  (syntax-rules()
    ((unless cond? expr ...)
     (if (not cond?) (begin expr ...)))))

;(unless (= x 0) (display x) (display " != 0") (newline))

(define-syntax for
  (syntax-rules (as in) 
    ((for index as where body ...)
     (for-each (lambda (where) body ...) index))
    ((for where in index body ...)
     (for index  as where body ...))))


;(for i in '(1 2 3)
;  (for j in '(4 5 6)
;    (display (list i j))
;    (newline))) 

;(for '(1 2 3) as i
;  (for '(4 5 6) as j
;    (display (list i j))
;    (newline)))


(define-syntax while
  (syntax-rules ()
    ((while cond? expr ...)
     (letrec ((iter (lambda ()
                      (if cond?
                          (begin
                            expr ...
                            (iter))))))
       (iter)))))

;(let ((p 0)
;      (q 0))
;  (while (< p 3)
;         (set! q 0)
;         (while (< q 3)
;                (display (list p q))
;                (newline)
;                (set! q (+ q 1)))
;         (set! p (+ p 1))))


(define-syntax repeat
  (syntax-rules (until)
    ((repeat (expr ...) until cond?)
     (letrec ((iter (lambda()
                      (begin
                        expr ...
                        (if (not cond?) (iter))))))
       (iter)))))

;(let ((i 0)
;      (j 0))
;  (repeat ((set! j 0)
;           (repeat ((display (list i j))
;                    (set! j (+ j 1)))
;                   until (= j 3))
;           (set! i (+ i 1))
;           (newline))
;          until (= i 3)))


(define-syntax cout
  (syntax-rules (« endl)
    ((_ « endl) (newline))
    ((_ « expr) (display expr))
    ((_ « endl « expr ...) (begin (newline) (cout « expr ...)))
    ((_ « expr1 « expr2 ...)
     (begin (display expr1)
            (cout « expr2 ...)))))