import './App.css';
import {useState} from "react";

function App() {
    const [keyword, setKeyword] = useState("");
    const [groups, setGroups] = useState("");
    const [result, setResult] = useState(null);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        setKeyword("")
        setGroups("")
        let groupArray = groups.split(/[ \n]/)
        console.log(groupArray)
        let data = {
            text: keyword,
            links: groupArray
        }
        fetch("http://localhost:8080/groups", {
            method: "POST",
            body: JSON.stringify(data)
        })
            .then(response => response.json())
            .then(result => setResult(result))
    };

    const printDate = (time) => {
        let x = new Date(time * 1000)
        return x.toLocaleString("ru-RU")
    }

    return (
        <div className="App">
            <form onSubmit={handleSubmit} className="Form">
                <div className={'row'}>
                    <label htmlFor="keyword">Keyword:</label>
                    <input className={"keyword"} id={"keyword"} type="text" name="Keyword" value={keyword}
                           onChange={(e => {
                               setKeyword(e.target.value)
                           })} required={true}/>
                </div>
                <div>
                    <label htmlFor="textarea">Groups:</label>
                    <textarea className={"textarea"} id={"textarea"} name="Groups" value={groups} onChange={(e => {
                        setGroups(e.target.value)
                    })} required={true}/>
                </div>
                <div className={"row"}>
                    <input className={"button"} type="submit" value="Search"/>
                </div>
            </form>
            {result === null ? ":(" : result.map(item => {
                return (
                    <div>
                        <a href={"https://" + item.Link} target={"_blank"}>{item.Link}</a>
                        <p>{printDate(item.Date)}</p>
                        <p className={"text"}>{item.Text}</p>

                    </div>
                )
            })}
        </div>
    );
}

export default App;
