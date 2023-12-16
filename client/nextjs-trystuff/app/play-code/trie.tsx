
'use client';
import { FormEvent, useState } from 'react';
import { Button, Card, CardHeader, FormHelperText, FormControl, FormGroup, Input, InputLabel, TextField, Toolbar} from '@mui/material';
export default function Trie() {

    let [wordString, setWordString] = useState('');
    let [matchingWords, setMatchingWords] = useState([]);
    let [prefix, setPrefix] = useState('');
    const getWords = async() => {
        const wordList = wordString.split(',');
        const res = await
        fetch("http://localhost:8080/", {
            method: 'POST',
            body: JSON.stringify({words: wordList, prefix: prefix}),
            headers: {'Content-Type': 'application/json'} 
        });
        const wordsMatchingPrefix = await res.json();
        setMatchingWords(wordsMatchingPrefix.data);
    }

    var cardStyle = {
        display: 'block',
        transitionDuration: '0.3s',
    }
    var wordListstyle = {
        margin: '1vw'
    }

    return (
        <>
        <FormGroup>
            <ol>
                <h1>
                    <li>
                        Enter a word list
                    </li>
                </h1>
                <h1>
                    <li>
                        Enter a prefix
                    </li>
                </h1>
                <h1>
                    <li>
                        Click GET MATCHING WORDS
                    </li>
                </h1>
            </ol>

                
            <FormControl>
                <InputLabel htmlFor="wordlist"></InputLabel>
                <TextField data-testid="wordlist" type="text" id="wordlist" name="wordlist" variant="outlined"
                        value={wordString}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                          setWordString(event.target.value);
                        }}
                    placeholder="Apple, Orange, Lime" aria-describedby="wordlist-helper-text"  />
                <FormHelperText id="wordlist-helper-text">Comma Separated dropdown values</FormHelperText>
            </FormControl>
            <FormControl>    
                <InputLabel htmlFor="prefix"></InputLabel>
                <TextField data-testid="prefix" type="text" id="prefix" name="prefix" variant="outlined"
                        value={prefix}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                            setPrefix(event.target.value);
                        }} placeholder="any prefix string" aria-describedby="prefix-helper-text" />
                <FormHelperText id="prefix-helper-text">String prefix to filter list</FormHelperText>
                <Toolbar sx={{ justifyContent: "space-between" }}>
                    <Button data-testid="submit" variant="contained" type="submit" onClick={getWords}>Get Matching Words</Button>
                </Toolbar>
            </FormControl>
        </FormGroup>
        <Card style={cardStyle}>
        <CardHeader
            title="Matching Words"
        />
            <div data-testid="matching" style={wordListstyle}>
                {matchingWords.map(paragraph => <p>{paragraph}</p>)}
           </div>
        </Card>
        </>
  )
}
