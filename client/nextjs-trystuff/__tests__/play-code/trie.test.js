import { render, screen } from "@testing-library/react";
import Trie from '../../app/play-code/trie';
import '@testing-library/jest-dom';


//test block
test("Trie is correct", () => {
    // render the component on virtual dom
    render(<Trie />);
    
    //select the elements you want to interact with
    const wordlist = screen.getByTestId("wordlist");
    const prefix = screen.getByTestId("prefix");
    const submit = screen.getByTestId("submit");
    const matching = screen.getByTestId("matching");
    
    //assert the expected result
    expect(wordlist).toBeInTheDocument();
    expect(prefix).toBeInTheDocument();
    expect(submit).toBeInTheDocument();
    expect(submit).toHaveTextContent("Get Matching Words");
    expect(matching).toBeInTheDocument();
});