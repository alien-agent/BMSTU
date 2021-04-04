package com.transparentjeans;

import java.util.Arrays;
import java.util.Optional;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class VectorSequence {
    public final int length;
    public final Vector[] vectors;

    public VectorSequence(int count) {
        this.length = count;
        this.vectors = IntStream.range(0, count).mapToObj(i -> new Vector()).toArray(Vector[]::new);
    }

    @Override
    public String toString() {
        return String.format("VectorSequence[%d]:\n", this.length) +
                Arrays.stream(this.vectors).map(Vector::toString).collect(Collectors.joining(", "));
    }

    public static Vector crossProduct(Vector v1, Vector v2) {
        int resultX = v1.z * v2.y - v1.y * v2.z,
                resultY = v1.x * v2.z - v1.z * v2.x,
                resultZ = v1.y * v2.x - v1.x * v2.y;
        return new Vector(resultX, resultY, resultZ);
    }

    public Stream<Vector> crossProductsStream(){
        return IntStream.range(1, this.length).mapToObj(i -> crossProduct(this.vectors[i-1], this.vectors[i]));
    }

    public Optional<Vector> findLongest(){
        return Arrays.stream(this.vectors).max(Vector::compareTo);
    }
}
