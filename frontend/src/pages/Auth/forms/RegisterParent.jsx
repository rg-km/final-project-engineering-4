import { Button, Flex } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import InputPassword from 'src/components/InputPassword';
import InputText from 'src/components/InputText';

const RegisterParentForm = () => {
  const initilInput = {
    nama: '',
    username: '',
    email: '',
    jenis_kelamin: '',
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

  const onSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    console.log(input);
    setIsLoading(false);
  };

  return (
    <form onSubmit={onSubmit}>
      <InputText name="nama" label="Nama" value={input.nama} handleChange={handleChange} isRequired />
      <InputText name="username" label="Username" value={input.username} handleChange={handleChange} isRequired />
      <InputText name="email" label="Email" value={input.email} handleChange={handleChange} isRequired />
      <InputText
        name="jenis_kelamin"
        label="Jenis Kelamin"
        value={input.jenis_kelamin}
        handleChange={handleChange}
        isRequired
      />
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
  );
};

export default RegisterParentForm;
