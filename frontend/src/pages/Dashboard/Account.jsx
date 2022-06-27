import { ChevronDownIcon } from "@chakra-ui/icons";
import { Box, Button, Flex, Heading, Image, Menu, MenuButton, MenuItem, MenuList, Stack, Text, VStack } from "@chakra-ui/react";
import Navbar from "./components/Navbar";

const Account = () => {
    return <Box bg={'gray.100'} minH={'100vh'}>
        <Navbar />
        <Box px={{ "base": 2, "md": 20 }} py={{ "base": 4, "md": 8 }}>
            <Box
                w={{ base: '100%', md: '50%' }}
                shadow={'sm'}
                m={'auto'}
                rounded={'md'}
                bg={'white'}
                p={{ "base": 2, "md": 4 }}>
                <Heading fontSize={{ base: "xs", md: 'sm' }}>Akun</Heading>
                <Box py={{ base: 4 }}>
                    <Image
                        borderRadius='full'
                        w={'30%'}
                        m={'auto'}
                        src='https://bit.ly/dan-abramov'
                        alt='Dan Abramov'
                    />
                    <Box textAlign={'center'} my={{ base: 2 }}>
                        <Menu>
                            <MenuButton size={'xs'} as={Button} variant={'none'} rightIcon={<ChevronDownIcon />}>
                                Ganti Profile
                            </MenuButton>
                            <MenuList fontSize={{base: 'xs', md: 'md'}}>
                                <MenuItem>Hapus Profile</MenuItem>
                                <MenuItem>Upload Baru</MenuItem>
                            </MenuList>
                        </Menu>
                    </Box>
                    <VStack spacing={2} align='stretch'>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>Nama</Text>
                            <Text fontSize={'sm'}>John Doe</Text>
                        </Flex>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>Role</Text>
                            <Text fontSize={'sm'}>Guru</Text>
                        </Flex>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>Email</Text>
                            <Text fontSize={'sm'}>johndoe@gmail.com</Text>
                        </Flex>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>Jenis Kelamin</Text>
                            <Text fontSize={'sm'}>Laki-laki</Text>
                        </Flex>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>Alamat</Text>
                            <Text fontSize={'sm'}>Semarang, Jawa Tengah</Text>
                        </Flex>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>No. HP</Text>
                            <Text fontSize={'sm'}>08563736662</Text>
                        </Flex>
                        <Flex justifyContent={'space-between'}>
                            <Text fontSize={'sm'}>Pendidikan</Text>
                            <Text fontSize={'sm'}>Sarjana Komputer</Text>
                        </Flex>
                    </VStack>
                    <Stack mt={{ base: 4 }} direction={{ base: 'column', md: 'row-reverse' }}>
                        <Button size={{ base: 'xs', md: 'sm' }}>Edit</Button>
                        <Button size={{ base: 'xs', md: 'sm' }} variant={'outline'}>Hapus Akun</Button>
                    </Stack>
                </Box>
            </Box>
        </Box>
    </Box>
}

export default Account;