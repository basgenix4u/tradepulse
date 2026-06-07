import { useState } from 'react'

function App() {
  const [user, setUser] = useState(null)
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const register = async () => {
    const res = await fetch('https://tradepulse-backend-u62m.onrender.com/api/v1/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password, name: 'User' })
    })
    const data = await res.json()
    if (data.token) {
      setUser(data.user)
      localStorage.setItem('token', data.token)
    }
  }

  const login = async () => {
    const res = await fetch('https://tradepulse-backend-u62m.onrender.com/api/v1/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    })
    const data = await res.json()
    if (data.token) {
      setUser(data.user)
      localStorage.setItem('token', data.token)
    }
  }

  if (user) {
    return (
      <div className="min-h-screen bg-[#0B1120] text-white p-8">
        <h1 className="text-3xl font-bold mb-4">Welcome {user.name}!</h1>
        <p className="text-gray-400">TradePulse is working!</p>
        <p className="mt-4">Email: {user.email}</p>
        <button onClick={() => { setUser(null); localStorage.clear() }} className="mt-4 px-4 py-2 bg-red-600 rounded">
          Logout
        </button>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-[#0B1120] flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        <h1 className="text-4xl font-bold text-center mb-8 text-cyan-400">TradePulse</h1>
        <div className="bg-[#111827] p-8 rounded-2xl border border-white/10">
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            className="w-full px-4 py-3 mb-4 bg-[#0B1120] border border-white/10 rounded-lg text-white"
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            className="w-full px-4 py-3 mb-6 bg-[#0B1120] border border-white/10 rounded-lg text-white"
          />
          <button onClick={register} className="w-full py-3 mb-3 bg-cyan-500 text-black font-semibold rounded-lg">
            Register
          </button>
          <button onClick={login} className="w-full py-3 bg-white/10 text-white font-semibold rounded-lg">
            Login
          </button>
        </div>
      </div>
    </div>
  )
}

export default App
