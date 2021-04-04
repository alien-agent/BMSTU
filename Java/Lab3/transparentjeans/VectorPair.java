package com.transparentjeans;

public class VectorPair implements Comparable<VectorPair> {
    public final Vector first;
    public final Vector second;
    public final Vector crossProduct;
    
    public VectorPair(Vector first, Vector second){
        this.first = first;
        this.second = second;
        this.crossProduct = this.calculateCrossProduct();
    }
    
    private Vector calculateCrossProduct(){
        int new_x = this.first.z * this.second.y - this.first.y * this.second.z,
                new_y = this.first.x * this.second.z - this.first.z * this.second.x,
                new_z = this.first.y * this.second.x - this.first.x * this.second.y;
        return new Vector(new_x, new_y, new_z);
    }

    public String toString(){
        return String.format("%-20s X %20s [%.2f]", this.first, this.second, this.crossProduct.length);
    }

    @Override
    public int compareTo(VectorPair other) {
        return Double.compare(this.crossProduct.length, other.crossProduct.length);
    }
}
