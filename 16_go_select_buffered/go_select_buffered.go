package main

import (
	"fmt"
	"time"
)

/*

	Όταν θέλουμε να έχουμε περιπτώσεις επιλογών σε go routines η Go έχει μια ειδική
	εντολή που ονομάζεται "select" η οποία λειτουργεί σαν την switch
	αλλά για τα κανάλια

	Το παρακάτω πρόγραμμα εμφανίζει "Από 1" και "Απο 3". Η select
	παίρνει το πρώτο κανάλι που είναι έτοιμο και λαμβάνει από αυτό δεδομένα
	(ή στέλνει σε αυτό). Αν περισσότερα από ένα κανάλια είναι έτοιμα
	τότε επιλέγει τυχαία ένα από το οποίο θα λαμβάνει. Εάν κανένα
	από τα κανάλια 	δεν είναι έτοιμο, τότε η διαδικασία σταματάει μέχρι
	κάποιο να γίνει διαθέσιμο.

*/

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 3)
	/*
		Είναι επίσης δυνατόν να περάσουμε μια δεύτερη παράμετρο στη
		συνάρτηση make, που είδαμε όταν δημιουργούμε ένα κανάλι
		Το πρώτο δημιουργεί ένα buffered κανάλι με χωρητικότητα 1 και το δεύτερο
		με χωρητικότητα 3.
		Κανονικά τα κανάλια είναι συγχρονισμένα δηλαδή οι δύο πλευρές του
		καναλιού θα περιμένουν έως ότου η άλλη πλευρά να είναι έτοιμη.
		Ένα buffered κανάλι απο την άλλη, είναι ασύγχρονο δηλαδή στην αποστολή
		ή στη λήψη ενός μηνύματος δεν θα περιμένει, εκτός αν το κανάλι
		είναι ήδη πλήρες.
	*/

	go func() {
		for {
			c1 <- "Απο 1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "Απο 3"
			time.Sleep(time.Second * 3)

		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Είσαι καλά;", msg1)
			case msg2 := <-c2:
				fmt.Println("Είμαι μια χαρά", msg2)
			/* Η εντολή select χρησιμοποιείται συχνά για
			να εφαρμόσει ένα timeout.
			Το time.After δημιουργεί ένα κανάλι και μετά τη δοθείσα
			διάρκεια, θα στείλει τον τρέχον χρόνο σε αυτό.
			(εμείς δεν ενδιαφερόμαστε για το χρόνο γιαυτό δεν τον
			αποθηκεύουμε σε μια μεταβλητή)*/
			case <-time.After(time.Second):
				fmt.Println("Λήξη χρόνου")
			default: // Η default case θα συμβεί αμέσως αν
				// κανένα από τα κανάλια δεν είναι έτοιμα.
				fmt.Println("Κανένα έτοιμο")
				//time.Sleep(time.Millisecond * 11000)
			}
			time.Sleep(time.Millisecond * 11000)
		}
	}()

	var input string
	fmt.Scanln(&input) // Περιμένει να πατήσουμε ENTER για να τερματίσει το πρόγραμμα
}
