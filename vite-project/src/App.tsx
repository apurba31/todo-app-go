import {Box} from '@mantine/core';
import './App.css';
import useSWR from 'swr';
export const ENDPOINT = 'http://localhost:4000'

const fetcher = (url: string) => 
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
  // const [count, setCount] = useState(0)
  const { data, mutate}  = useSWR('api/todos', fetcher)

  return (
    <div>
      <h1>JSON.stringify(data)</h1>
      {/* <Box>Hello World</Box> */}
    </div>
    
  )
}

export default App
