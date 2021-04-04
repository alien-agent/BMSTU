package com.transparentjeans;

import java.util.Iterator;
import java.util.Random;

public class Polynom implements Iterable<Polynom> {
    private final int[] coefficients;
    public int degree;

    public Polynom() {
        Random rnd = new Random();
        this.coefficients = new int[rnd.nextInt(5) + 5];
        for (int i = 0; i < this.coefficients.length; i++) {
            this.coefficients[i] = rnd.nextInt(40) - 20;
        }
        this.degree = this.coefficients.length - 1;
    }

    public Polynom(int[] coefficients) {
        this.coefficients = coefficients;
        this.degree = coefficients.length - 1;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        int coef;
        for (int i = 0; i <= this.degree; i++) {
            // Add coefficient
            if ((coef = this.coefficients[i]) != 0) {
                if (coef > 0) {
                    if (i != 0) {
                        sb.append(" + ");
                    }
                    if (coef != 1) {
                        sb.append(coef);
                    }
                } else {
                    sb.append(" - ");
                    if (coef != -1) {
                        sb.append(-coef);
                    }
                }
            }
            // Add "x" with power if needed
            if (this.degree - i > 1) {
                sb.append(String.format("x^%s", this.degree - i));
            } else if (this.degree - i > 0) {
                sb.append("x");
            }
        }

        return sb.toString();
    }

    public Polynom derivative() {
        int[] newCoefficients = new int[this.degree];
        for (int i = 0; i < this.degree; i++) {
            newCoefficients[i] = this.coefficients[i] * (this.degree - i);
        }
        return new Polynom(newCoefficients);
    }

    @Override
    public Iterator<Polynom> iterator() {
        return new Iterator<>() {
            private Polynom current = new Polynom(coefficients);

            @Override
            public boolean hasNext() {
                return current.degree > 0;
            }

            @Override
            public Polynom next() {
                current = current.derivative();
                return current;
            }
        };
    }
}
