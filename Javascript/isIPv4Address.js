// Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.

function isIPv4Address(inputString) {

    var currentNumber = 0;
    var digits = 0;
    var countNumbers = 0;
  
    inputString += '.';
  
    for (var i = 0; i < inputString.length; i++) {
      if (inputString[i] === '.') {
        if (digits == 0) {
          return false;
        }
        if (currentNumber.toString().length !== digits) {
          return false;
        }
        countNumbers++;currentNumber=0;digits=0;
      }
      else {
        var digit = inputString.charCodeAt(i) - '0'.charCodeAt(0);
        if (digit < 0 || digit > 9) {
          return false;
        }
        digits++;
        currentNumber = currentNumber * 10 + digit;
        if (currentNumber > 255) {
          return false;
        }
      }
    }
    return countNumbers === 4;
  }
  