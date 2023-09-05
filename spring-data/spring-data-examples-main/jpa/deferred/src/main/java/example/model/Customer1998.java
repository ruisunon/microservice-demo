package example.model;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class Customer1998 {

	@Id @GeneratedValue(strategy = GenerationType.AUTO) private long id;
	private String firstName;
	private String lastName;

	protected Customer1998() {}

	public Customer1998(String firstName, String lastName) {
		this.firstName = firstName;
		this.lastName = lastName;
	}

	@Override
	public String toString() {
		return String.format("Customer1998[id=%d, firstName='%s', lastName='%s']", id, firstName, lastName);
	}

}