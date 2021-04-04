package com.transparentjeans;

import java.util.Arrays;

public class Test {
    private static void printArray(VectorPair[] array) {
        StringBuilder result = new StringBuilder();
        result.append(String.format("\n%12s %25s %12s\n", "Vector1", "Vector2", "Length"));

        for (int i = 0; i < array.length; i++) {
            result.append(array[i].toString()).append("\n");
        }

        System.out.println(result);
    }

    public static void main(String[] args) {
        int vectorsCount = 10;
        VectorPair[] pairs = new VectorPair[vectorsCount];

        for (int i = 0; i < vectorsCount; i++) {
            pairs[i] = new VectorPair(new Vector(), new Vector());
        }

        System.out.println("Random generated array:");
        printArray(pairs);
        Arrays.sort(pairs);
        System.out.println("Array after sorting:");
        printArray(pairs);
    }
}