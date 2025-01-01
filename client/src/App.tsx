// import { useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'

import { BrowserRouter } from 'react-router';
import { RouterWrapper } from './components';
import './App.css';

function App() {
    return (
        <BrowserRouter>
            <RouterWrapper />
        </BrowserRouter>
    );
}

export default App;
