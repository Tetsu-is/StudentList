"use client";

import { FC, Suspense, useEffect, useState } from "react";

const Loading: FC = () => {
  return <div>Loading...</div>;
};

type Student = {
  student_id: number;
  name: string;
  age: number;
};

const page = () => {
  const [students, setStudents] = useState<Student[]>([]);

  const fetchStudents = async () => {
    const res = await fetch('http://localhost:8080/student');
    const data = await res.json();
    setStudents(data);
  }
 
  return (
    <>
      <h1>Students</h1>
      <table>
        <thead>
          <tr>
            <th scope="col">StudentID</th>
            <th scope="col">Name</th>
            <th scope="col">Age</th>
          </tr>
        </thead>
        <tbody>
          {students.map((student)=> {
            return (
                <tr key = {student.student_id}>
                    <th scope="row">
                        {student.student_id}
                    </th>
                    <td>
                        {student.name}
                    </td>
                    <td>
                        {student.age}
                    </td>
                </tr>
            )
          })}
        </tbody>
      </table>
      <button type="button" onClick={fetchStudents}>Get Students</button>
    </>
  );
};

export default page;
