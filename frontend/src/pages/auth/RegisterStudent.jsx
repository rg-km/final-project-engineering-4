import { Box, Button, Flex, Heading, Text } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AppAlert from 'src/components/AppAlert';
import BgAuth from 'src/components/BgAuth';
import InputPassword from 'src/components/InputPassword';
import InputText from 'src/components/InputText';

const RegisterStudent = () => {
    const initilInput = {
        nama: "",
        username: "",
        email: "",
        jenis_kelamin: "",
        tempat: "",
        tanggal_lahir: "",
        no_hp: "",
        alamat: "",
        password: ""
    }
    const [isLoading, setIsLoading] = useState(false);
    const [input, setInput] = useState(initilInput);
    const [alert, setAlert] = useState({
        status: false,
        type: "success",
        message: ""
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
    }

    return <BgAuth>
        <Box overflow={'auto'} bg={'white'} maxH={'80vh'} w={['90%', '50%']} m={'auto'} p={[4, 6, 10, 20]} borderRadius={['xl', '2xl']} boxShadow={['2xl']}>
            <Heading textAlign={'center'} fontSize={['sm', 'md', 'lg']} mb={[3, 4, 8]}>Daftar sebagai Siswa</Heading>

            <AppAlert alert={alert} setAlert={setAlert} />
            
            <form onSubmit={onSubmit}>
                <InputText name='nama' label='Nama' value={input.nama} handleChange={handleChange} isRequired />
                <InputText name='username' label='Username' value={input.username} handleChange={handleChange} isRequired />
                <InputText name='email' label='Email' value={input.email} handleChange={handleChange} isRequired />
                <InputText name='jenis_kelamin' label='Jenis Kelamin' value={input.jenis_kelamin} handleChange={handleChange} isRequired />
                <InputText name='tempat' label='Tempat Lahir' value={input.tempat} handleChange={handleChange} />
                <InputText name='tanggal_lahir' label='Tanggal Lahir' value={input.tanggal_lahir} handleChange={handleChange} />
                <InputText name='no_hp' label='No HP' value={input.no_hp} handleChange={handleChange} isRequired />
                <InputText name='alamat' label='Alamat' value={input.alamat} handleChange={handleChange} />
                <InputPassword name='password' label='Password' value={input.password} handleChange={handleChange} isRequired />

                <Flex justify={'center'} gap={[2]} mb={[2, 4]}>
                    <Link to={'/register'}><Button size={['sm', 'md']}>Kembali</Button></Link>
                    <Button type='submit' isLoading={isLoading} colorScheme='blue' size={['sm', 'md']}>Daftar</Button>
                </Flex>
            </form>

            <Text textAlign={'center'} fontSize={['xs', 'sm']}>Atau <Link to={'/login'} style={{ textDecoration: 'underline' }}>masuk</Link></Text>
        </Box>
    </BgAuth>
}

export default RegisterStudent;