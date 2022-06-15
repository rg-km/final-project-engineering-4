import { BrowserRouter, Route, Routes } from "react-router-dom";
import Register from "src/pages/auth/Register";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/register" element={<Register />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
