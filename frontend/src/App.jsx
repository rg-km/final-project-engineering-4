import { Route, Routes } from 'react-router-dom';
import Login from '@routes';
import LoginParent from '@routes';
import LoginStudent from '@routes';
import LoginTeacher from '@routes';
import Register from '@routes';
import RegisterParent from '@routes';
import RegisterStudent from '@routes';
import RegisterTeacher from '@routes';

const App = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Register />} />
        <Route path="/register" element={<Register />} />
        <Route path="/register/teacher" element={<RegisterTeacher />} />
        <Route path="/register/parent" element={<RegisterParent />} />
        <Route path="/register/student" element={<RegisterStudent />} />
        <Route path="/login" element={<Login />} />
        <Route path="/login/teacher" element={<LoginTeacher />} />
        <Route path="/login/parent" element={<LoginParent />} />
        <Route path="/login/student" element={<LoginStudent />} />
      </Routes>
    </div>
  );
};

export default App;
