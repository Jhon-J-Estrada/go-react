import React from 'react'

function App() {
  return (
    <div>

        <h1>Hola desde react</h1>    
        <button onClick={async () =>{
            const respon = await fetch('http://127.0.0.1:3000/users')
            const data = await respon.json()

            console.log(data)
        }}>
          Obtener Datos
        </button>

    </div>

    
  )
}

export default App
