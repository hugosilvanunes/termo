import {
  Component,
  createEffect,
  createResource,
  createSignal,
  JSX,
} from "solid-js";
import "./app.css";
import Row from "./Row";
import Rows from "./Rows";

const fetchGame = async () =>
  (await fetch(import.meta.env.VITE_APP_API_URL)).json();

interface Game {
  length: number;
}

const App: Component = () => {
  const [attempts, setAttempts] = createSignal<number>(5);
  const [length, setLength] = createSignal<number>(0);
  const [data] = createResource<Game>(fetchGame);

  createEffect(() => {
    const wordLength = data()?.length;

    if (wordLength) setLength(wordLength);
  });

  const getValue = () => {
    console.log("chegou");
  };

  return (
    <div
      style={{
        display: "flex",
        "flex-direction": "column",
        "align-items": "center",
      }}
    >
      <header style={{ margin: "10px" }}>
        Palavra com {length} caracteres
      </header>
      <div style={{ "margin-top": "50px" }}>
        <form onSubmit={getValue}>
          <Rows wordLength={length()} attemptLength={attempts()} />
        </form>
      </div>
    </div>
  );
};

export default App;
