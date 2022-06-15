import { Box, Flex, Heading, Text } from '@chakra-ui/react';
import React from 'react';
import { Link } from 'react-router-dom';
import RoleItem from 'src/components/RoleItem';
import bg_auth from "@images/bg_auth.jpg";
import teacher from "@images/teaching.png";

const Register = () => {
    return <Box bg={'gray.50'} h={'100vh'} display={'flex'} bgImg={bg_auth} bgRepeat={'no-repeat'} bgSize={'cover'}>
        <Box bg={'white'} w={['90%', '50%']} m={'auto'} p={[4, 6, 10, 20]} borderRadius={['xl', '2xl']} boxShadow={['2xl']}>
            <Heading textAlign={'center'} fontSize={['sm', 'md', 'lg']} mb={[3, 4, 8]}>Daftar sebagai</Heading>
            <Flex gap={[2, 4, 8]} justifyContent={'center'} mb={[2, 4]}>
                <RoleItem title={'Guru'} redirect={'/register/teacher'} img={teacher} bg={'yellow.200'} />
                <RoleItem title={'Orang Tua'} redirect={'/register/parent'} img={teacher} bg={'blue.200'} />
                <RoleItem title={'Siswa'} redirect={'/register/student'} img={teacher} bg={'green.200'} />
            </Flex>
            <Text textAlign={'center'} fontSize={['xs', 'sm']}>Atau <Link to={'/login'} style={{ textDecoration: 'underline' }}>masuk</Link></Text>
        </Box>
    </Box>
}

export default Register;