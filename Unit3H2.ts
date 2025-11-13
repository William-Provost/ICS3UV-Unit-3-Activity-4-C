/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-13
 * @fileoverview This program tells you what to pay for a movie.
 */

// variables
let ageAsString: string = "";
let ageAsNumber: number = 0;

// input age
ageAsString = prompt("Please enter your age: ") || "0";
ageAsNumber = parseInt(ageAsString);

// check the age
if (ageAsNumber >= 65) {
  console.log("You are a senior");
  console.log("Please pay $15.00 to see the show");
} else if (ageAsNumber >= 19) {
  console.log("You are an adult");
  console.log("Please pay $25.00 to see the show");
} else if (ageAsNumber >= 12) {
  console.log("You are a teenager");
  console.log("Please pay $20.00 to see the show");
} else if (ageAsNumber >= 6) {
  console.log("You are a child");
  console.log("Please pay $15.00 to see the show");
} else {
  console.log("You are not in school");
  console.log("You are free to see the show");
}

console.log("\nDone.");