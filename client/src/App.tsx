import { useEffect, useState } from "react";

type note = {
  title: string
}

function App() {
  const [notes, setNotes] = useState<note[]>([]);
  useEffect(() => {
    fetch("/api/notes")
      .then(res =>
        res.json()
      )
      .then(body =>
        setNotes(body)
      )
  }, [])
  return <>
    <button onClick={() => {
      fetch("/api/notes?title=foobar", { method: "POST" })
    }}>POST</button>
    <ul>
      {notes.map((note, i) => <li id={note.title + i}>{note.title}</li>)}
    </ul>
  </>;
}

export default App;
