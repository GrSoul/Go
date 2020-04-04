package main

import (
	"fmt"
	"time"
)

// GO ROUTINES

/*
Ένας διακομιστής ιστοσελίδων λαμβάνει αιτήσεις που υποβάλλονται από προγράμματα
περιήγησης των χρηστών στο διαδίκτυο και παρέχει HTML ιστοσελίδες ως απάντηση.
Κάθε αίτημα λαμβάνεται σαν ένα μικρό πρόγραμμα.

Θα ήταν ιδανικό για προγράμματα όπως αυτά να είναι σε θέση να τρέξουν μικρότερα
στοιχεία τους την ίδια χρονική στιγμή (στην περίπτωση του διακομιστή ιστοσελίδων
να λαμβάνει πολλαπλές αιτήσεις).Η επίτευξη προόδου σε περισσότερες από μία
εργασίες ταυτόχρονα είναι γνωστό ως ταυτοχρονισμός. Η Go έχει εγγενής υποστήριξη
για τον ταυτοχρονισμό χρησιμοποιώντας go-ρουτίνες και κανάλια channels

*/

// Φτιάχνουμε μια ρουτίνα εργασίας που μετράει απο το 0 έως το 10
func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)

		// Παύση της λειτουργίας για 1 δευτερόλεπτο για να επιτραπεί
		// η εκτέλεση άλλων λειτουργιών
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {

	// Μια go ρουτίνα είναι μια func που τρέχει ταυτόχρονα με άλλες func
	// Ορίζουμε παρακάτω μια τέτοια, χρησιμοποιώντας την λέξη κλειδί "go"
	// ακολουθούμενη από το όνομα της func που θα τρέξει.
	// Εδω λοιπόν η go-ρουτίνα είναι μία συνάρτηση που είναι ικανή να εκτελεί
	// ταυτόχρονα και άλλες συναρτήσεις (την count(i)).

	for i := 0; i < 10; i++ {
		go count(i)
	}

	// Ας προσθέσουμε κάποια καθυστέρηση στη συνάρτηση με το time.Sleep
	// για να βεβαιωθούμε ότι η go ρουτίνα έχει χρόνο για να τελειώσει,
	// ειδάλως το πρόγραμμα θα τελειώσει πριν συμβεί αυτό
	time.Sleep(time.Millisecond * 11000)
}

// Σημείωση: Όταν τρέξετε το πρόγραμμα η count εμφανίζει τους αριθμούς από 0 έως 10, με αναμονή 1 δευτερολέπτου μετά
// από κάθε εκτέλεση. Οι go-ρουτίνα με αυτό τον τρόπο εκτελείται ταυτόχρονα.
