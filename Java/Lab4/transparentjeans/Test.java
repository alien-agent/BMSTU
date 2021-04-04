package com.transparentjeans;

public class Test {

    public static void main(String[] args) {
        Polynom p = new Polynom();
        for(Polynom currentDerivative:p){
            System.out.println(currentDerivative);
        }
    }
}