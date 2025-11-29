/**
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-28
 * @fileoverview  User enters 5 rows and receives correct number triangle.
 */

const rowString: string = prompt("How many rows would you like? ") || "0";
const rows: number = parseInt(rowString);

for (let i: number = 1; i <= rows; i++) {
  let line: string = "";
  for (let j: number = 1; j <= i; j++) {
    line += j + " ";
  }
  console.log(line);
}
console.log("\nDone.");
