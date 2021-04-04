package com.transparentjeans;

import java.util.stream.Stream;

import static java.lang.Math.pow;
import static java.lang.Math.sqrt;

public class Vector {
    public final int x, y, z;
    public final double length;

    public Vector() {
        this(rnd(), rnd(), rnd());
    }

    public Vector(int x, int y, int z) {
        this.x = x;
        this.y = y;
        this.z = z;
        this.length = this.calculateLength();
    }

    private static int rnd() {
        return (int) ((Math.random() * (40)) - 20);
    }

    public String toString() {
        return String.format("Vector(%s,%s,%s)", this.x, this.y, this.z);
    }

    /*public int compareTo(Vector other) {
        return Double.compare(this.length, other.length);
    }*/

    private double calculateLength() {
        int sum = Stream.of(this.x, this.y, this.z).mapToInt(x -> (int) pow(x, 2)).sum();
        return sqrt(sum);
    }
}