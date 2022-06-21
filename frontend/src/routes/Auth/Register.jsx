import { useState } from 'react';
import { Link } from 'react-router-dom';
import { Button, Stack, Text } from '@chakra-ui/react';
import { PATH } from '@routes/path';
import AuthContainer from './components/AuthContainer';
import InputText from '@components/InputText';
import InputPassword from '@components/InputPassword';

const Register = ({ role }) => {
  const initilInput = {
    nama: '',
    username: '',
    email: '',
    jenis_kelamin: '',
    pendidikan: '',
    agama: '',
    no_hp: '',
    alamat: '',
    password: '',
  };

  const [isLoading, setIsLoading] = useState(false);
  const [input, setInput] = useState(initilInput);

  const handleChange = (e) => {
    setInput({
      ...input,
      [e.target.name]: e.target.value,
    });
  };

  const onSubmit = (e) => {
    e.preventDefault();
    setIsLoading(true);
    console.log(input);
    setIsLoading(false);
  };

  return (
    <AuthContainer
      title={`Daftar sebagai ${role.title}`}
      subtitle="Lengkapi form di bawah dengan menggunakan data Anda yang valid">
      <Stack as={'form'} onSubmit={onSubmit} spacing={8}>
        <Stack spacing={3}>
          <InputText name="nama" label="Nama" value={input.nama} handleChange={handleChange} isRequired />
          <InputText
            name="username"
            label="Username"
            value={input.username}
            handleChange={handleChange}
            isRequired
          />
          <InputText name="email" label="Email" value={input.email} handleChange={handleChange} isRequired />
          <InputText
            name="jenis_kelamin"
            label="Jenis Kelamin"
            value={input.jenis_kelamin}
            handleChange={handleChange}
            isRequired
          />
          <InputText name="pendidikan" label="Pendidikan" value={input.pendidikan} handleChange={handleChange} />
          <InputText name="agama" label="Agama" value={input.agama} handleChange={handleChange} />
          <InputText name="no_hp" label="No HP" value={input.no_hp} handleChange={handleChange} isRequired />
          <InputText name="alamat" label="Alamat" value={input.alamat} handleChange={handleChange} />
          <InputPassword
            name="password"
            label="Password"
            value={input.password}
            handleChange={handleChange}
            isRequired
          />
        </Stack>

        <Button type="submit" isLoading={isLoading}>
          Daftar
        </Button>
      </Stack>

      <Stack>
        <Text textAlign={'center'} fontSize={'sm'}>
          Sudah punya akun?{' '}
          <Button as={Link} size={'sm'} variant={'link'} to={`${PATH.LOGIN}/${role.value}`}>
            Masuk sebagai {role.value}
          </Button>
        </Text>
      </Stack>
    </AuthContainer>
  );
};

export default Register;
