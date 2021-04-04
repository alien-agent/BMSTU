package com.transparentjeans;

import java.util.Optional;
import java.util.stream.Collectors;

public class Test {
    static final int lengthThreshold = 150; // По заданию должно быть равно 1, но для наглядности лучше ~150

    public static void main(String[] args) {
        VectorSequence sequence = new VectorSequence(7);
        System.out.println(sequence);

        var result = sequence.crossProductsStream()
                .collect(Collectors.groupingBy(v -> v.length > lengthThreshold));
        System.out.printf("\nCross products by length:\n<%s: %s\n>%s: %s\n",
                lengthThreshold, result.get(false), lengthThreshold, result.get(true));

        Optional<Vector> longest = sequence.findLongest();
        System.out.printf("\nLongest: %s", longest.isPresent() ? longest.get() : "missing");
    }
}