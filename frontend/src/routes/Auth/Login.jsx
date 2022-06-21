import { Button, Flex } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AppAlert from '@components/AppAlert';
import InputPassword from '@components/InputPassword';
import InputText from '@components/InputText';
import AuthFormContainer from './components/AuthFormContainer';

const Login = ({ role }) => {
  const initilInput = {
    email: '',
    password: '',
  };
  const [isLoading, setIsLoading] = useState(false);
  const [input, setInput] = useState(initilInput);
  const [alert, setAlert] = useState({
    status: false,
    type: 'warning',
    message: '',
  });

  const handleChange = (e) => {
    setInput({
      ...input,
      [e.target.name]: e.target.value,
    });
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    console.log(input);
    setIsLoading(false);
  };

  return (
    <AuthFormContainer title={`Masuk sebagai ${role.title}`}>
      <AppAlert alert={alert} setAlert={setAlert} />

      <form onSubmit={onSubmit}>
        <InputText name="email" label="Email" value={input.email} handleChange={handleChange} isRequired />
        <InputPassword
          name="password"
          label="Password"
          value={input.password}
          handleChange={handleChange}
          isRequired
        />

        <Flex justify={'center'} gap={[2]}>
          <Link to={'/login'}>
            <Button size={['sm', 'md']}>Kembali</Button>
          </Link>
          <Button type="submit" isLoading={isLoading} colorScheme="blue" size={['sm', 'md']}>
            Login
          </Button>
        </Flex>
      </form>
    </AuthFormContainer>
  );
};

export default Login;
