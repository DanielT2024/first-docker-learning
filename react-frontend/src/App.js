import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [count, setCount] = useState(0);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCount = async () => {
      try {
        // IMPORTANT: Use the Docker host IP or 'localhost' if mapping ports
        // When running React in its container and Go in its container,
        // React will call the Go backend via the mapped port on the Docker host.
        // In a production Kubernetes scenario, you'd use a service name.
        // For this Docker-only scenario, we use 'localhost' and the exposed port.
        const response = await fetch('http://localhost:8000/api/count'); 

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setCount(data.count);
      } catch (e) {
        console.error("Failed to fetch count:", e);
        setError("Failed to load counter. Backend might not be running or accessible.");
      }
    };

    fetchCount();
  }, []); // Empty dependency array means this runs once on mount

  return (
    <div className="App">
      <header className="App-header">
        <h1>Simple React App with Go Backend & Redis</h1>
        {error ? (
          <p style={{ color: 'red' }}>{error}</p>
        ) : (
          <p>Page views: {count}</p>
        )}
        <p>
          This React app talks to a Go backend, which talks to a Redis database, 
          all running in separate Docker containers!
        </p>
      </header>
    </div>
  );
}

export default App;