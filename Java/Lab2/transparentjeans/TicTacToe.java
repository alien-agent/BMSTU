package com.transparentjeans;

import java.util.Arrays;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class TicTacToe {
    private final String[][] board;
    public int size;
    public boolean isCrossNext = true; // Первым ходит крестик
    public boolean isGameOver = false;
    public String winner = null;
    private int cellsOccupied = 0;

    public TicTacToe(int size) {
        if (size <= 0) {
            throw new IllegalArgumentException("Размер доски должен быть положительным!");
        }
        this.size = size;
        this.board = new String[size][size];
    }

    private static boolean isWinningLine(Stream<String> line) {
        var unique = line.distinct().toArray(String[]::new);
        return unique.length == 1 && unique[0] != null;
    }

    private static String centerString(int width, String s) {
        return String.format("%-" + (s.length() + (width - s.length()) / 2) + "s", String.format("%" + width + "s", s));
    }

    @Override
    public String toString() {
        StringBuilder result = new StringBuilder();

        for (int i = 0; i < this.size; i++) {
            result.append("|");
            for (int j = 0; j < this.size; j++) {
                result.append(centerString(4, this.board[i][j] == null ? "*" : this.board[i][j]));
            }
            result.append("|\n");
        }

        return result.toString().trim();
    }

    private void checkGameOver(int row, int col) {
        Stream<String> rowStream = Arrays.stream(this.board[row]),
                colStream = Arrays.stream(this.board).map(r -> r[col]),
                mainDiagonalStream = IntStream.range(0, this.size).mapToObj(i -> this.board[i][i]),
                sideDiagonalStream = IntStream.range(0, this.size).mapToObj(i -> this.board[this.size - 1 - i][i]);

        if (isWinningLine(rowStream) ||
                isWinningLine(colStream) ||
                isWinningLine(mainDiagonalStream) ||
                isWinningLine(sideDiagonalStream)) {
            this.isGameOver = true;
            this.winner = this.board[row][col];
        } else if (this.cellsOccupied == this.size * this.size) {
            this.isGameOver = true;
            this.winner = "дружба";
        }
    }

    public void move(int row, int col) {
        if (this.isGameOver) {
            throw new IllegalArgumentException("Невозможно совершить ход: игра окончена!");
        } else if (row >= this.size || col >= this.size || row < 0 || col < 0) {
            throw new IllegalArgumentException("Невозможно совершить ход: координаты выходят за границы доски");
        } else if (this.board[row][col] != null) {
            throw new IllegalArgumentException("Невозможно совершить ход: значение уже установлено");
        } else {
            this.board[row][col] = this.isCrossNext ? "✕" : "○";
            this.cellsOccupied++;
            this.isCrossNext = !this.isCrossNext;
            this.checkGameOver(row, col);
        }
    }
}
