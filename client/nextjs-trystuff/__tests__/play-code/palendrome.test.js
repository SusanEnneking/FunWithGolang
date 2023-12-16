import { render, screen } from "@testing-library/react";
import Palendrome from '../../app/play-code/palendrome';
import '@testing-library/jest-dom';


//test block
test("Palendrome is correct", () => {
    // render the component on virtual dom
    render(<Palendrome />);
    
    //select the elements you want to interact with
    const palendrome = screen.getByTestId("palendrome");

    
    //assert the expected result
    expect(palendrome).toHaveTextContent("tacocat");
});