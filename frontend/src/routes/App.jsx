import { BrowserRouter, Route, Routes } from "react-router-dom";
import Register from "src/pages/auth/Register";
import RegisterTeacher from "src/pages/auth/RegisterTeacher";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/register" element={<Register />} />
        <Route path="/register/teacher" element={<RegisterTeacher />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
