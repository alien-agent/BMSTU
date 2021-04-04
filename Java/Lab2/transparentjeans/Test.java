package com.transparentjeans;

import java.util.Scanner;

public class Test {
    public static void clearScreen() {
        System.out.print("\033[H\033[2J");
        System.out.flush();
    }

    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);
        System.out.print("Введите размер доски: ");
        int size = scanner.nextInt();
        TicTacToe game = new TicTacToe(size);

        System.out.println("Чтобы сделать ход, введите через пробел номер строки и столбца(нумерация с 0).");
        while(!game.isGameOver){
            System.out.print("Сейчас ходит " + (game.isCrossNext ? "✕" : "○") + ": ");
            int row = scanner.nextInt(), col = scanner.nextInt();
            try {
                game.move(row, col);
                System.out.println(game);
            } catch (Exception e){
                System.out.println(e.getMessage());
            }
        }
        String winMessage = game.winner.equals("дружба") ? "Победила " : "Победил ";
        System.out.println(winMessage + game.winner + "!");
    }
}
