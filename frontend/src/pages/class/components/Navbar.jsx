import { ChevronDownIcon } from "@chakra-ui/icons";
import { Box, Button, Flex, Menu, MenuButton, MenuItem, MenuList } from "@chakra-ui/react";
import { PATH } from "@config/path";
import { APP } from "@utils/constants";
import { Link, useNavigate } from "react-router-dom";

const Navbar = () => {
    let navigate = useNavigate();
    let currClass = {
        id_class: "qwerty",
        nama_kelas: "9F",
        kode_kelas: "qwertyasdf",
        keterangan: "Kelas dengan wali kelas yang bernama Bu Nurul Huda",
    }
    return <Flex bg={'white'} alignItems={'center'} justifyContent={'space-between'} boxShadow={{ base: 'sm' }} py={{ base: 1, md: 2 }} px={{ base: 2, md: 20 }}>
        <Box>
            <Button as={Link} to={PATH.HOME} variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }}>{APP.name}</Button>
        </Box>
        <Flex gap={{ "base": 2, "md": 6 }}>
            <Button variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }}>{currClass.nama_kelas}</Button>
            <Menu>
                <MenuButton as={Button} variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }} rightIcon={<ChevronDownIcon />}>
                    John Doe
                </MenuButton>
                <MenuList fontSize={{ base: 'xs', md: 'md' }}>
                    <MenuItem onClick={() => navigate(PATH.DASHBOARD)}>Home</MenuItem>
                    <MenuItem onClick={() => navigate(PATH.ACCOUNT)}>Akun</MenuItem>
                </MenuList>
            </Menu>
        </Flex>
    </Flex>
}

export default Navbar;