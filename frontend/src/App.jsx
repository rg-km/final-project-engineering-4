import { Route, Routes } from 'react-router-dom';
import Home from '@routes/Home/Home';
import Register from '@routes/Auth/Register';

const App = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </div>
  );
};

export default App;
