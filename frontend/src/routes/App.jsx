import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "src/pages/auth/Login";
import LoginParent from "src/pages/auth/LoginParent";
import LoginStudent from "src/pages/auth/LoginStudent";
import LoginTeacher from "src/pages/auth/LoginTeacher";
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
        <Route path="/login" element={<Login />} />
        <Route path="/login/teacher" element={<LoginTeacher />} />
        <Route path="/login/parent" element={<LoginParent />} />
        <Route path="/login/student" element={<LoginStudent />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
