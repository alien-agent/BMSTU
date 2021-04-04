(load "unit-test.rkt")
(define (derive_difference expr)
  (cond ((= (length expr) 2) (- (derivative (cadr expr))))
        ((= (length expr) 3) (list '- (derivative (cadr expr)) (derivative (caddr expr))))))

(define (derive_sum expr)
  (if (= (length expr) 3)
      
      (list '+ (derivative (cadr expr))
            (derivative (caddr expr)))
      
      (list '+ (derivative (cadr expr))
            (derivative (cons '+ (cddr expr))))))


(define (derive_product expr)
  (if (and (number? (cadr expr)) (equal? 'x (caddr expr)))
      (cadr expr)
      (if (= (length expr) 3)
          (list '+ (list '* (derivative (cadr expr)) (caddr expr)) (list '* (cadr expr) (derivative (caddr expr))))
          (list '+ (list '* (derivative (cadr expr)) (cons '* (cddr expr))) (list '* (cadr expr) (derivative (cons '* (cddr expr))))))))
      
(define (derive_division expr)
  (list '/ (list '- (list '* (derivative (cadr expr)) (caddr expr))
                 (list '* (cadr expr) (derivative (caddr expr))))
        (list 'expt (caddr expr) 2)))

(define (derivative expr)
  (cond
    ((number? expr) '0) ; константа
    ((equal? expr 'x) '1) ; переменная
    ((equal? (car expr) '-) (derive_difference expr)) ; разность
    ((equal? (car expr) '+) (derive_sum expr)) ; сумма
    ((equal? (car expr) '*) (derive_product expr)) ; произведение
    ((equal? (car expr) '/) (derive_division expr)) ; деление
    ((and (equal? (car expr) 'expt) (number? (caddr expr))) ; x^a
     (list '* (caddr expr) (list 'expt (cadr expr) (- (caddr expr) 1))
           (derivative (cadr expr))))            
    ((and (equal? (car expr) 'expt) (list? (caddr expr)) ; x^(f(x))
          (equal? '- (caaddr expr)) (number? (car (cdr (cadr (cdr expr))))))

     (list '* (caddr expr) (list 'expt (cadr expr) (list '- (caddr expr) 1))
           (derivative (cadr expr))))
    ((and (equal? (car expr) 'expt) (number? (caddr expr))) (list '* expr (list 'log (cadr expr)) (derivative (caddr expr)))) ; a^x
    ((equal? (car expr) 'exp) (list '* expr (derivative (cadr expr)))) ; e^x
    ((equal? (car expr) 'cos) (list '* (list '- (list 'sin (cadr expr))) (derivative (cadr expr)))) ; cosx
    ((equal? (car expr) 'sin) (list '* (list 'cos (cadr expr)) (derivative (cadr expr)))) ; sinx
    ((equal? (car expr) 'log) (list '* (list 'expt (cadr expr) (- 1)) (derivative (cadr expr)))) ; lnx
    ;возводим х в степень
    ((equal? (car expr) 'expt) (list '* expr (list 'log (cadr expr)) (derivative (caddr expr))))))

(define many-many-tests
  (list (test (derivative 2) 0)
        (test (derivative 'x) 1)
        (test (derivative '(- x)) -1)
        (test (derivative '(* 1 x)) 1)
        (test (derivative '(* -1 x)) -1)
        (test (derivative '(* -4 x)) -4)
        (test (derivative '(* 10 x)) 10)
        (test (derivative '(- (* 2 x) 3)) '(- 2 0))
        (test (derivative '(* 2 (expt x 5))) '(+ (* 0 (expt x 5)) (* 2 (* 5 (expt x 4) 1))))
        (test (derivative '(expt x 10)) '(* 10 (expt x 9) 1))
        (test (derivative '(expt x -2)) '(* -2 (expt x -3) 1))
        (test (derivative '(expt 5 x)) '(* (expt 5 x) (log 5) 1))
        (test (derivative '(cos x)) '(* (- (sin x)) 1))
        (test (derivative '(sin x)) '(* (cos x) 1))
        (test (derivative '(exp x)) '(* (exp x) 1))
        (test (derivative '(* 2 (exp x))) '(+ (* 0 (exp x)) (* 2 (* (exp x) 1))))
        (test (derivative '(log x)) '(* (expt x -1) 1))
        (test (derivative '(+ (expt x 3) (expt x 2))) '(+ (* 3 (expt x 2) 1) (* 2 (expt x 1) 1)))
        (test (derivative '(/ 3 x)) '(/ (- (* 0 x) (* 3 1)) (expt x 2)))
        (test (derivative '(* 2 (sin x) (cos x))) '(+ (* 0 (* (sin x) (cos x))) (* 2 (+ (* (* (cos x) 1) (cos x)) (* (sin x) (* (- (sin x)) 1))))))
        (test (derivative '(sin (log (expt x 2)))) '(* (cos (log (expt x 2))) (* (expt (expt x 2) -1) (* 2 (expt x 1) 1))))))


(display "----- Testing HOMEWORK -----")
(newline)
(run-tests many-many-tests)