import { Component } from "solid-js";
import Row from "./Row";

interface RowsProps {
  attemptLength: number;
  wordLength: number;
}

const Rows: Component = (props: RowsProps) => {
  return (
    <>
      {[...Array(props.attemptLength)].map((_, i) => (
        <Row length={props.wordLength} />
      ))}
    </>
  );
};

export default Rows;
