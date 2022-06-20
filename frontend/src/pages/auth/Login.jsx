import { Box, Flex, Heading, Text } from '@chakra-ui/react';
import React from 'react';
import BgAuth from 'src/components/BgAuth';
import RoleItem from 'src/components/RoleItem';
import teacher from "@images/teaching.png";
import parent from "@images/parents.png";
import student from "@images/student.png";
import { Link } from 'react-router-dom';

const Login = () => {
    return <BgAuth>
        <Box bg={'white'} w={['90%', '50%']} m={'auto'} p={[4, 6, 10, 20]} borderRadius={['xl', '2xl']} boxShadow={['2xl']}>
            <Heading textAlign={'center'} fontSize={['sm', 'md', 'lg']} mb={[3, 4, 8]}>Masuk sebagai</Heading>
            <Flex gap={[2, 4, 8]} justifyContent={'center'} mb={[2, 4]}>
                <RoleItem title={'Guru'} redirect={'/login/teacher'} img={teacher} bg={'yellow.200'} />
                <RoleItem title={'Orang Tua'} redirect={'/login/parent'} img={parent} bg={'blue.200'} />
                <RoleItem title={'Siswa'} redirect={'/login/student'} img={student} bg={'green.200'} />
            </Flex>
            <Text textAlign={'center'} fontSize={['xs', 'sm']}>Belum punya akun? <Link to={'/register'} style={{ textDecoration: 'underline' }}>daftar</Link></Text>
        </Box>
    </BgAuth>
}

export default Login;