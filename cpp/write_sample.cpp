// $ clang++ cpp/write_sample.cpp cpp/person.pb.cc -lprotobuf
// $ ./a.out

#include <iostream>
#include <fstream>
#include <string>

#include "person.pb.h"

int main(void) {
    GOOGLE_PROTOBUF_VERIFY_VERSION;

    Person person;
    person.set_id(123);
    person.set_name("NAME from C++");
    person.set_email("cpp@example.com");

    std::ofstream fout("sample_output_data.pb", std::ios::binary);
    person.SerializeToOstream(&fout);

    return 0;
}
