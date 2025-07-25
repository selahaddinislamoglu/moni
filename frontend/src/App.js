import React, { useEffect, useState } from 'react';
import { LineChart, Line, XAxis, YAxis, Tooltip, CartesianGrid, ResponsiveContainer } from 'recharts';

function App() {
  const [data, setData] = useState([]);

  // Fetch CPU usage every second
  useEffect(() => {
    const interval = setInterval(async () => {
      try {
        const res = await fetch('http://localhost:8080/cpu/usage/last-five-seconds');
        const json = await res.json();

        // Add new point, limit to last 10 readings
        setData(prev => {
          const newData = [...prev, {
            time: new Date(json.time * 1000).toLocaleTimeString(),
            usage: json.usage,
          }];
          return newData.slice(-100);
        });
      } catch (err) {
        console.error('Fetch error:', err);
      }
    }, 5000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div style={{ width: '80%', height: '400px', padding: '20px' }}>
      <h2>Real-time CPU Usage</h2>
      <ResponsiveContainer width="100%" height="100%">
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="time" />
          <YAxis domain={[0, 100]} unit="%" />
          <Tooltip />
          <Line type="monotone" dataKey="usage" stroke="#8884d8" dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}

export default App;
