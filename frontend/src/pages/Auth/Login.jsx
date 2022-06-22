import { Button, Stack, Text } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AppAlert from '@components/AppAlert';
import InputPassword from '@components/InputPassword';
import InputText from '@components/InputText';
import AuthContainer from './components/AuthContainer';
import { PATH } from '@config/path';

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
    <AuthContainer title={`Masuk sebagai ${role.title}`} subtitle="Masukkan email dan password Anda untuk masuk.">
      <AppAlert alert={alert} setAlert={setAlert} />

      <Stack as={'form'} onSubmit={onSubmit} spacing={8}>
        <Stack spacing={3}>
          <InputText name="email" label="Email" value={input.email} handleChange={handleChange} isRequired />
          <InputPassword
            name="password"
            label="Password"
            value={input.password}
            handleChange={handleChange}
            isRequired
          />
        </Stack>

        <Stack direction={'row'}>
          <Button type="submit" isLoading={isLoading} flex={1}>
            Login
          </Button>
        </Stack>
      </Stack>
      <Stack>
        <Text textAlign={'center'} fontSize={'sm'}>
          Belum punya akun?{' '}
          <Button as={Link} size={'sm'} variant={'link'} to={`${PATH.REGISTER}/${role.value}`}>
            Daftar sebagai {role.value}
          </Button>
        </Text>
      </Stack>
    </AuthContainer>
  );
};

export default Login;
