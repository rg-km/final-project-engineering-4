import { CheckIcon, ChevronLeftIcon, ChevronRightIcon, CloseIcon, DeleteIcon } from "@chakra-ui/icons";
import { Box, Button, Checkbox, Flex, IconButton, Input, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, Stack, Table, TableContainer, Tbody, Td, Th, Thead, Tr, useDisclosure, VStack } from "@chakra-ui/react";
import { PATH } from "@config/path";
// import { useState } from "react";
import Card from "./components/Card";
import MenuItem from "./components/MenuItem";
import Navbar from "./components/Navbar";

const Class = () => {
    const { isOpen, onOpen, onClose } = useDisclosure();

    const handleSubmit = () => {
        // console.log(input);
    }

    return <Box bg={'gray.100'} minH={'100vh'}>
        <Navbar />
        <Box px={{ base: 2, md: 20 }} py={{ base: 2, md: 4 }}>
            <Stack direction={{ base: "column", md: "row" }} spacing={{ base: 4 }}>
                <VStack w={"20%"} spacing={{ base: 2, md: 4 }} align='stretch'>
                    <MenuItem label={"Absen"} redirect={PATH.CLASS + "qwerty"} />
                    <MenuItem label={"Jadwal"} redirect={PATH.CLASS + "qwerty" + "/schedules"} />
                    <MenuItem label={"Siswa"} redirect={PATH.CLASS + "qwerty" + "/students"} />
                    <MenuItem label={"Orang Tua"} redirect={PATH.CLASS + "qwerty" + "/parents"} />
                    <MenuItem label={"Pengaturan"} redirect={PATH.CLASS + "qwerty" + "/setting"} />
                </VStack>
                <Box>
                    <Card title={"Absen"}>
                        <Flex justifyContent={"space-between"}>
                            <Button size={{ base: "xs", md: "sm" }} onClick={onOpen}>Buat Absen</Button>
                            <Input size={{ base: "xs", md: "sm" }} type="text" placeholder="25-06-2022" w={"auto"} />
                        </Flex>
                        <Box textAlign={"right"}>
                            <IconButton variant={'unstyled'} icon={<ChevronLeftIcon />} size={"sm"} />
                            <IconButton variant={'unstyled'} icon={<ChevronRightIcon />} size={"sm"} />
                        </Box>
                        <TableContainer>
                            <Table size={{ base: "sm" }} overflow={"auto"}>
                                <Thead>
                                    <Tr>
                                        <Th>Nama</Th>
                                        <Th>20/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                        <Th>21/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                        <Th>22/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                        <Th>23/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                        <Th>24/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                        <Th>25/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                        <Th>26/06/2022<IconButton variant={'link'} color={"red.500"} icon={<DeleteIcon />} size={"xs"} /></Th>
                                    </Tr>
                                </Thead>
                                <Tbody>
                                    <Tr>
                                        <Td>Rahman Maulana</Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"red.500"} icon={<CloseIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                    </Tr>
                                    <Tr>
                                        <Td>Jordan Maulana</Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"red.500"} icon={<CloseIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                    </Tr>
                                    <Tr>
                                        <Td>Ragil Maulana</Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"red.500"} icon={<CloseIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                    </Tr>
                                    <Tr>
                                        <Td>Bayu Maulana</Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"red.500"} icon={<CloseIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                    </Tr>
                                    <Tr>
                                        <Td>Hasan Maulana</Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"red.500"} icon={<CloseIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                    </Tr>
                                    <Tr>
                                        <Td>Basuki Maulana</Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"red.500"} icon={<CloseIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                        <Td><IconButton variant={'unstyled'} color={"green.500"} icon={<CheckIcon />} size={"sm"} /></Td>
                                    </Tr>
                                </Tbody>
                            </Table>
                        </TableContainer>
                    </Card>
                </Box>
            </Stack>

            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader fontSize={{ base: 'md' }}>Buat Absen</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <VStack align={"stretch"} spacing={2}>
                            <Checkbox value={"1"} fontSize={{base: "xs", md: "sm"}}>Rahman Maulana</Checkbox>
                            <Checkbox value={"1"} fontSize={{base: "xs", md: "sm"}}>Jordan Maulana</Checkbox>
                            <Checkbox value={"1"} fontSize={{base: "xs", md: "sm"}}>Ragil Maulana</Checkbox>
                            <Checkbox value={"1"} fontSize={{base: "xs", md: "sm"}}>Bayu Maulana</Checkbox>
                            <Checkbox value={"1"} fontSize={{base: "xs", md: "sm"}}>Hasan Maulana</Checkbox>
                            <Checkbox value={"1"} fontSize={{base: "xs", md: "sm"}}>Basuki Maulana</Checkbox>
                        </VStack>
                    </ModalBody>
                    <ModalFooter>
                        <Button colorScheme='blue' onClick={handleSubmit} size={{ base: 'sm' }}>
                            Simpan
                        </Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </Box>
    </Box>
}

export default Class;