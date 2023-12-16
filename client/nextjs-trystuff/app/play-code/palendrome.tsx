
import styles from '../page.module.css'
export default function Palendrome() {
    

  const palindrome = (strText: string) => {
    let firstHalf = '';
    let firstHalfArray = [];
    let secondHalf = '';
    let centerIndex = Math.floor(strText.length/2);
    if (strText.length % 2 != 0){
        firstHalf = strText.substring(0, centerIndex + 1);
    } else {
        firstHalf = strText.substring(0, centerIndex);
    }
    firstHalfArray = firstHalf.split('');
    firstHalf = firstHalfArray.reverse().join('');
    secondHalf = strText.substring(centerIndex, strText.length);
    if (firstHalf === secondHalf){
        return strText;
    }

    return '';
  }
  const palindromeCheck = (strText: string) =>{
    let longest = '';
    for (let center = 0; center < strText.length; center++){
        for (let plusMinus = 0; plusMinus < strText.length; plusMinus++){
            if (center + plusMinus <= strText.length && center - plusMinus > -1){
                let newLongest = palindrome(strText.substring(center - plusMinus, center + plusMinus + 1));
                if (longest.length < newLongest.length){
                    longest = newLongest;
                }
            }
        }
    }
    return longest;
  }

  const output = palindromeCheck("tacocat");
  return (
    <div>
        <h1>Problem Output</h1>
        <div  className={styles.card}>
            <p data-testid="palendrome">{output}</p>
        </div>
    </div>
  )
}
