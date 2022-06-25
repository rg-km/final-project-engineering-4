import { Box, Button, Heading, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, SimpleGrid, useDisclosure } from "@chakra-ui/react";
import InputText from "@components/InputText";
import { PATH } from "@config/path";
import { useEffect, useState } from "react";
import AddClass from "./components/AddClass";
import ClassItem from "./components/ClassItem";
import Navbar from "./components/Navbar";

const initialInput = {
    nama_kelas: '',
    keterangan: ''
}

const Dashboard = () => {
    let data = [
        {
            id_class: 1,
            nama_kelas: "9F",
            kode_kelas: "qwertyasdf",
            keterangan: "Kelas dengan wali kelas yang bernama Bu Nurul Huda",
        },
        {
            id_class: 2,
            nama_kelas: "9E",
            kode_kelas: "qwertyuiop",
            keterangan: "Kelas dengan wali kelas yang bernama Bu Sum",
        },
        {
            id_class: 3,
            nama_kelas: "9D",
            kode_kelas: "zxcvasdf",
            keterangan: "Kelas dengan wali kelas yang bernama Pak Rizky",
        }
    ];

    const [classes, setClasses] = useState([]);
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [input, setInput] = useState(initialInput);

    useEffect(() => {
        setClasses(data);
    }, [])

    const handleSubmit = () => {
        console.log(input);
    }

    return <Box bg={'gray.100'} minH={'100vh'}>
        <Navbar />
        <Box px={{ "base": 2, "md": 20 }} py={{ "base": 4, "md": 8 }}>
            <Heading fontSize={{ "base": "xs", "md": "md" }} mb={{ "base": 2 }}>Daftar Kelas Anda</Heading>

            <SimpleGrid columns={{ "base": 2, "md": 4 }} spacing={{ "base": 2, md: 4 }}>
                <AddClass to={PATH.DASHBOARD} onClick={onOpen} />
                {
                    classes.length > 0 && classes.map((kelas, index) => {
                        return <ClassItem key={index} to={PATH.DASHBOARD} kelas={kelas} />
                    })
                }
            </SimpleGrid>

            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader fontSize={{base: 'md'}}>Tambah Kelas Baru</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <InputText type={'text'} label={'Nama Kelas'} value={input.nama_kelas} onChange={(e) => setInput({ ...input, nama_kelas: e.target.value })} required />
                        <InputText type={'text'} label={'Keterangan'} value={input.keterangan} onChange={(e) => setInput({ ...input, keterangan: e.target.value })} />
                    </ModalBody>
                    <ModalFooter>
                        <Button colorScheme='blue' onClick={handleSubmit} size={{base: 'sm'}}>
                            Tambah Baru
                        </Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </Box>
    </Box>
}

export default Dashboard;