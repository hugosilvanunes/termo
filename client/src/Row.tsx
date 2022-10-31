import { Component } from "solid-js";

interface RowProps {
  length?: number;
}

const Row: Component = (props: RowProps) => {
  return (
    <div
      style={{
        display: "flex",
        "flex-direction": "row",
      }}
    >
      {[...Array(props.length)].map((_, i) => (
        <input
          style={{
            width: "50px",
            height: "60px",
            margin: "10px",
          }}
          type="text"
          maxLength="1"
          id={`${i}`}
        />
      ))}
    </div>
  );
};

export default Row;
