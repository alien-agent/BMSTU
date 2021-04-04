

; Task 1


(define-syntax trace-ex
  (syntax-rules ()
    ((_ expr)
     (begin
       (write 'expr)
       (display " => ")
       (let ((result expr))
         (write result)
         (newline)
         result)))))

(define (zip . xss)
  (if (or (null? xss)
          (null? (trace-ex (car xss))))
      '()
      (cons (map car xss)
            (apply zip (map cdr (trace-ex xss))))))

(display "----- Testing task 1 -----")
(newline)
(zip '(1 2 3) '(one two three))

; Task 2

(define (run-test the-test)
  (let ((expr (car the-test)))
    (write expr)
    (let* ((result (eval expr (interaction-environment)))
           (status (equal? (cadr the-test) result)))
      (if status
          (begin (display " OK") (newline))
          (begin (display " FAIL")
                 (newline)
                 (display "  Expected: ") (write (cadr the-test))
                 (newline)
                 (display "  Returned: ") (write result) (newline))) status)))

(define (run-tests tests)
  (define (and-fold x xs)
    (if (null? xs) x (and-fold (and x (car xs)) (cdr xs))))
  (and-fold #t (map run-test tests)))

(define-syntax test
  (syntax-rules ()
    ((_ expr expected-result) (list 'expr expected-result))))


(define (signum x)
  (cond
    ((< x 0) -1)
    ((= x 0)  1) ; Ошибка здесь!
    (else     1)))

(define tests2
  (list (test (signum -2) -1)
        (test (signum  0)  0)
        (test (signum  2)  1)))

(display "----- Testing task 2 -----")
(newline)
(run-tests tests2)

; Task 3

(define (ref xs index . elem)
  ; Узнаем тип последовательности
  (define (get-type xs)
    (cond ((vector? xs) 'vector)
          ((string? xs) 'string)
          (else 'list)))
  ; Получение элемента по индексу
  (define (ref-get xs index)
    (if (eq? index 0)
        (car xs)
        (if (not (null? (cdr xs)))
            (ref-get (cdr xs) (- index 1))
            (not (null? (cdr xs))))))
  ; Вставка элемента
  (define (ref-insert xs index new)
    (if (eq? index 0)
        (append (append new (list (car elem))) xs)
        (if (not (null? (cdr xs)))
            (ref-insert (cdr xs) (- index 1) (append new (list (car xs))))
            #f)))
  ; Превращаем любую последовательность в список :)
  (define (translate xs)
    (cond ((vector? xs) (vector->list xs))
          ((string? xs) (string->list xs))
          (else xs)))
  ; Название говорит само за себя
  (define (list-of-char? xs status) 
    (if (null? xs)
        status
        (list-of-char? (cdr xs) (and (char? (car xs)) status))))
  ; Возвращаем последовательность в исходную форму :)
  (define (re-translate xs type)
    (cond ((eq? type 'vector) (list->vector xs))
          ((eq? type 'string) 
           (if (list-of-char? xs #t)
               (list->string xs)
               #f))
          (else xs)))
  ; Превращаем последовательность в список, совершаем нужное действие, и при необходимости возвращаем последовательность в исходную форму
  (if (null? elem)
      (ref-get (translate xs) index)
      (let ((temp (ref-insert (translate xs) index '())))
        (if (not temp) temp (re-translate temp (get-type xs))))))
      

(define tests3
  (list (test (ref '(1 2 3) 1) 2)
        (test (ref #(1 2 3) 1) 2)
        (test (ref '(1 2 3) 1 0) '(1 0 2 3))
        (test (ref #(1 2 3) 1 0) #(1 0 2 3))
        (test (ref "123" 1 #\0) "1023")
        (test (ref "123" 1 0) #f)
        (test (ref "123" 5 #\4) #f)))
(display "----- Testing task 3 -----")
(newline)
(run-tests tests3)

; Task 4

(define (factorize expr)
  (let ((a (cadr (cadr expr))) (b (cadr (caddr expr))))
    (cond
      ((and (equal? (caddr (cadr expr)) 2) (equal? (car expr) '-))
       (list '* (list '- a b) (list '+ a b)))
      ((and (equal? (caddr (cadr expr)) 3) (equal? (car expr) '-))
       (list '* (list '- a b) (list '+ `(expt a 2) (list '* a b) `(expt b 2))))
      ((and (equal? (caddr (cadr expr)) 3) (equal? (car expr) '+))
       (list '* (list '+ a b) (list '+ `(expt a 2) (list '- (list '* a b)) `(expt b 2))))
      (else expr))))

(display "----- Testing task 4 -----")
(newline)

(factorize '(- (expt x 2) (expt y 2))) 
; ⇒ (* (- x y) (+ x y))

(factorize '(- (expt (+ car 1) 2) (expt (- cadr 1) 2)))
; ⇒ (* (- (+ car 1) (- cadr 1)) (+ (+ car 1) (- cadr 1)))
         
(eval (list (list 'lambda 
                  '(x y) 
                  (factorize '(- (expt x 2) (expt y 2))))
            1 2) (interaction-environment))
; ------------------- Homework -----------------------
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