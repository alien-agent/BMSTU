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