package com.rydata.springboot3.services;

import java.util.List;
import java.util.Map;

import com.rydata.springboot3.repositories.EmployeeRepository;
import jakarta.persistence.EntityManager;
import jakarta.persistence.PersistenceContext;
import jakarta.persistence.StoredProcedureQuery;
import jakarta.transaction.Transactional;

import com.rydata.springboot3.Pojos.EmpIdName;
import com.rydata.springboot3.entities.Employee;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class EmployeeService {
    
    @Autowired
    EmployeeRepository employeeRepository;

    @PersistenceContext
    EntityManager em;

    public EmployeeService(){

    }

    public List<Employee> saveAllEmployees(List<Employee> employees){
        return employeeRepository.saveAll(employees);
    }

    public Employee saveEmployee(Employee employee){
        return employeeRepository.save(employee);
    }

    public List<Employee> findAllEmployees(){
        return employeeRepository.findAll();
    }

    public List<Employee> findAllEmployeesByIds(List<Integer> ids){
        return employeeRepository.findAllById(ids);
    }

    public Employee findEmployeeById(int id){
        return employeeRepository.findById(id);
    }

    public List<Employee> findEmployeesByName(String employeename){
        return employeeRepository.findByEmployeename(employeename);
    }

    public int Gettotalemployees(){
        return employeeRepository.ProcgetTotalEmployees();
    }

    public int getEmployeeAge(int employeeid){
        return employeeRepository.ProcGetEmployeeAge(employeeid);
    }

    public Map<String,?> getEmpNameandDept(int employeeid){
        return employeeRepository.ProcGetNameandDept(employeeid);
    }

    public List<EmpIdName> getIdAndName(int employeeid){

        StoredProcedureQuery spq = em.createNamedStoredProcedureQuery("emp.GetEmpIdandName");
        spq.setParameter("employeeid", employeeid);
        spq.execute();
        return spq.getResultList();

    }

    @Transactional
    public List<Employee> getAllEmpRef(int employeeid){

        StoredProcedureQuery spq = em.createNamedStoredProcedureQuery("emp.GetAllEmployees");
        spq.setParameter(2, employeeid);
        spq.execute();
        return spq.getResultList();

    }

}
