import './App.css';
import {useEffect, useState} from "react";
import axios from "axios"

function App() {
  const [counter, setCounter] = useState(0);

  useEffect(() => {
      axios.get("/api/counter")
          .then(r => {
              setCounter(r.data.Value)
          })
          .catch(err => {
              console.error(err)
          })
  },[setCounter])

  const increment = () => {
      axios.post("/api/counter/increment",{Value:1})
          .then(r => {
              setCounter(r.data.Value)
          })
          .catch(err => {
              console.error(err)
          })
  }

    const decrement = () => {
        axios.post("/api/counter/decrement",{Value:1})
            .then(r => {
                setCounter(r.data.Value)
            })
            .catch(err => {
                console.error(err)
            })
    }
  return (
    <div className="App">
      <h1>Counter Value : {counter}</h1>
      <button onClick={increment}>++</button>
      <button onClick={decrement}>--</button>
    </div>
  );
}

export default App;
