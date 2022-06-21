import { Box, Button, Flex, Heading, Text } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AppAlert from 'src/components/AppAlert';
import AuthBackground from '@routes/Auth/components/AuthBackground';
import InputPassword from 'src/components/InputPassword';
import InputText from 'src/components/InputText';

const RegisterTeacher = () => {
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
  const [alert, setAlert] = useState({
    status: false,
    type: 'success',
    message: '',
  });

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
    <AuthBackground>
      <Box
        overflow={'auto'}
        bg={'white'}
        maxH={'80vh'}
        w={['90%', '50%']}
        m={'auto'}
        p={[4, 6, 10, 20]}
        borderRadius={['xl', '2xl']}
        boxShadow={['2xl']}>
        <Heading textAlign={'center'} fontSize={['sm', 'md', 'lg']} mb={[3, 4, 8]}>
          Daftar sebagai Guru
        </Heading>

        <AppAlert alert={alert} setAlert={setAlert} />

        <form onSubmit={onSubmit}>
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

          <Flex justify={'center'} gap={[2]} mb={[2, 4]}>
            <Link to={'/register'}>
              <Button size={['sm', 'md']}>Kembali</Button>
            </Link>
            <Button type="submit" isLoading={isLoading} colorScheme="blue" size={['sm', 'md']}>
              Daftar
            </Button>
          </Flex>
        </form>

        <Text textAlign={'center'} fontSize={['xs', 'sm']}>
          Atau{' '}
          <Link to={'/login'} style={{ textDecoration: 'underline' }}>
            masuk
          </Link>
        </Text>
      </Box>
    </AuthBackground>
  );
};

export default RegisterTeacher;
