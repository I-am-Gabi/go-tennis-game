/**
 * Código que simula uma partida de tênis com dois jogadores, 
 * seguindo as especificações passadas na disciplina de 
 * Programação Concorrente.
 */
package main

import (
	"fmt"
	"math/rand" 
	"time"
)
 
// score
var scoreA_points, scoreB_points int 
// game
var gameA_points, gameB_points int 
// set
var setA_points, setB_points int 
// match
var matchA_points, matchB_points int 

const P int = 2
const G int = 2
const S int = 2
const vantage int = 1

func playerA(c chan string) { 
    for {
        c <- "a"  
    }
}

func playerB(c chan string) { 
    for {
        c <- "b"    
    }
} 

func new_set() {
    scoreA_points, scoreB_points = 0, 0
    gameA_points, gameB_points = 0, 0
}


func game_manager(c chan string) {
    for { 
        player := <- c
 
        var player_win string

        result := "perdeu" 
        if rand.Float32() < 0.5 {
        	result = "ganhou"  
        	if player == "a" { 
                player_win = "a" 
        		scoreA_points++
        	} else {
				scoreB_points++
                player_win = "b"
        	}
        } else {
            if player == "a" {  
                player_win = "b"
                scoreB_points++
            } else {
                player_win = "a"
                scoreA_points++
            }
        }

        time.Sleep(time.Second * 1)  

    	fmt.Println("[player " + player + "]: " + result) 
    	fmt.Println("[score] \n A: ", scoreA_points, " \n B: ", scoreB_points)

    	if player_win == "a" {
    		if scoreA_points >= P && scoreA_points >= scoreB_points + vantage {
    			gameA_points++
    			scoreA_points, scoreB_points = 0, 0
    		}

    		if gameA_points >= G && gameA_points >= gameB_points + vantage {
    			setA_points++
                new_set()
    		}

            if setA_points >= S && setA_points >= setB_points + vantage {
                fmt.Println("JOGADOR A GANHOU") 
                return
            }
    	} 

    	if player_win == "b" {
    	 	if scoreB_points >= P && scoreB_points >= scoreA_points + vantage {
    			gameB_points++
    			scoreA_points, scoreB_points = 0, 0
    		}

    		if gameB_points >= G && gameB_points >= gameA_points + vantage {
	    		setB_points++
                new_set()
	    	}

            if setB_points >= S && setB_points >= setA_points + vantage {
                fmt.Println("JOGADOR B GANHOU") 
                return
            }
    	}

    	fmt.Println("[game] \n A: ", gameA_points, " \n B: ", gameB_points)
    	fmt.Println("[set] \n A: ", setA_points, " \n B: ", setB_points)
    }
}

func main() {
	fmt.Println("Game start")   

	var c chan string = make(chan string)
 
	rand.Seed(time.Now().UnixNano())
 
	go playerA(c)
	go playerB(c)

	go game_manager(c)
 
 	var input string
    fmt.Scanln(&input)   
}