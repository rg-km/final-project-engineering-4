import { BrowserRouter, Route, Routes } from "react-router-dom";
import Register from "src/pages/auth/Register";
import RegisterParent from "src/pages/auth/RegisterParent";
import RegisterStudent from "src/pages/auth/RegisterStudent";
import RegisterTeacher from "src/pages/auth/RegisterTeacher";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/register" element={<Register />} />
        <Route path="/register/teacher" element={<RegisterTeacher />} />
        <Route path="/register/parent" element={<RegisterParent />} />
        <Route path="/register/student" element={<RegisterStudent />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
