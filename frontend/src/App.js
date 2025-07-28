import React, { useEffect, useState } from 'react';
import {
  LineChart, Line, XAxis, YAxis, Tooltip, CartesianGrid, ResponsiveContainer
} from 'recharts';

function App() {
  const [cpuData, setCpuData] = useState([]);
  const [memData, setMemData] = useState([]);
  const [diskData, setDiskData] = useState([]);
  const [networkData, setNetworkData] = useState([]);

  useEffect(() => {
    const interval = setInterval(async () => {
      try {
        // Fetch CPU
        const cpuRes = await fetch('http://localhost:8080/cpu/usage/last-five-seconds');
        const cpuJson = await cpuRes.json();

        // Fetch Memory
        const memRes = await fetch('http://localhost:8080/memory/usage/all');
        const memJson = await memRes.json();

        // Fetch Disk
        const diskRes = await fetch('http://localhost:8080/disk/usage/last-five-seconds');
        const diskJson = await diskRes.json();

        // Fetch Network
        const networkRes = await fetch('http://localhost:8080/network/usage/last-five-seconds');
        const networkJson = await networkRes.json();

        setCpuData(prev => {
          const newData = [...prev, { time: new Date(cpuJson.time * 1000).toLocaleTimeString(), usage: cpuJson.usage }];
          return newData.slice(-60);
        });

        setMemData(prev => {
          const newData = [...prev, { time: new Date(memJson.time * 1000).toLocaleTimeString(), usage: memJson.usage }];
          return newData.slice(-60);
        });

        setDiskData(prev => {
          const newData = [...prev, { time: new Date(diskJson.time * 1000).toLocaleTimeString(), usage: diskJson.usage }];
          return newData.slice(-60);
        });

        setNetworkData(prev => {
          const newData = [...prev, { time: new Date(networkJson.time * 1000).toLocaleTimeString(), usage: networkJson.usage }];
          return newData.slice(-60);
        });

      } catch (err) {
        console.error('Fetch error:', err);
      }
    }, 5000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div style={{ width: '80%', margin: 'auto', paddingTop: '20px' }}>
      <h2 style={{ marginTop: '40px' }}>Real-time CPU Usage</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={cpuData}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="time" />
          <YAxis domain={[0, 100]} unit="%" />
          <Tooltip />
          <Line type="monotone" dataKey="usage" stroke="#8884d8" dot={false} />
        </LineChart>
      </ResponsiveContainer>

      <h2 style={{ marginTop: '40px' }}>Real-time Memory Usage</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={memData}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="time" />
          <YAxis domain={[0, 100]} unit="%" />
          <Tooltip />
          <Line type="monotone" dataKey="usage" stroke="#82ca9d" dot={false} />
        </LineChart>
      </ResponsiveContainer>

      <h2 style={{ marginTop: '40px' }}>Real-time Disk Usage</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={diskData}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="time" />
          <YAxis domain={[0, 100]} unit="MB" />
          <Tooltip />
          <Line type="monotone" dataKey="usage" stroke="#82ca9d" dot={false} />
        </LineChart>
      </ResponsiveContainer>

      <h2 style={{ marginTop: '40px' }}>Real-time Network Usage</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={networkData}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="time" />
          <YAxis domain={[0, 100]} unit="MB" />
          <Tooltip />
          <Line type="monotone" dataKey="usage" stroke="#82ca9d" dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}

export default App;
