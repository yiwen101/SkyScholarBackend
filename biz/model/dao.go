package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/yiwen101/SkyScholarBackend/biz/model/course"
)

var courseMap map[string]course.TreeNode
var tocRoot *course.TreeNode

const toc = `
1 INTRODUCTION
1.1 WHAT IS AN OPERATING SYSTEM?
1.1.1 The Operating System as an Extended Machine
1.1.2 The Operating System as a Resource Manager
1.2 HISTORY OF OPERATING SYSTEMS
1.2.1 The First Generation (1945-55) Vacuum Tubes and Plugboards
1.2.2 The Second Generation (1955-65) Transistors and Batch Systems
1.2.3 The Third Generation (1965-1980) ICs and Multiprogramming
1.2.4 The Fourth Generation (1980-Present) Personal Computers
1.2.5 Ontogeny Recapitulates Phytogeny
1.3 THE OPERATING SYSTEM ZOO
1.3.1 Mainframe Operating Systems
1.3.2 Server Operating Systems
1.3.3 Multiprocessor Operating Systems
1.3.4 Personal Computer Operating Systems
1.3.5 Real-Time Operating Systems
1.3.6 Embedded Operating Systems
1.3.7 Smart Card Operating Systems
1.4 COMPUTER HARDWARE REVIEW
1.4.1 Processors
1.4.2 Memory
1.4.3 I/O Devices
1.4.4 Buses
1.5 OPERATING SYSTEM CONCEPTS
1.5.1 Processes
1.5.3 Memory Management
1.5.4 Input/Output
1.5.5 Files
1.5.6 Security
1.5.8 Recycling of Concepts
1.6 SYSTEM CALLS
1.6.1 System Calls for Process Management
1.6.2 System Calls for File Management
1.6.3 System Calls for Directory Management
1.6.4 Miscellaneous System Calls
1.6.5 The Windows Win32 API
1.7 OPERATING SYSTEM STRUCTURE
1.7.1 Monolithic Systems
1.7.2 Layered Systems
1.7.3 Virtual Machines
1.7.4 Exokernels
1.7.5 Client-Server Model
1.8 RESEARCH ON OPERATING SYSTEMS
1.9 OUTLINE OF THE REST OF THIS BOOK
1.10 METRIC UNITS
1.11 SUMMARY
2 PROCESSES AND THREADS
2.1 PROCESSES
2.1.1 The Process Model
2.1.2 Process Creation
2.1.3 Process Termination
2.1.4 Process Hierarchies
2.1.5 Process States
2.1.6 Implementation of Processes
2.2 THREADS
2.2.1 The Thread Model
2.2.2 Thread Usage
2.2.3 Implementing Threads in User Space
2.2.4 Implementing Threads in the Kernel
2.2.5 Hybrid Implementations
2.2.6 Scheduler Activations
2.2.7 Pop-Up Threads
2.2.8 Making Single-Threaded Code Multithreaded
2.3 INTERPROCESS COMMUNICATION
2.3.1 Race Conditions
2.3.2 Critical Regions
2.3.3 Mutual Exclusion with Busy Waiting
Disabling Interrupts Lock Variables Strict Alternation Peterson’s Solution The TSL Instruction
2.3.4 Sleep and Wakeup
The Producer-Consumer Problem
2.3.5 Semaphores
Solving the Producer-Consumer Problem using Semaphores
2.3.6 Mutexes
2.3.7 Monitors
2.3.8 Message Passing
Design Issues for Message Passing Systems
The Producer-Consumer Problem with Message Passing
2.3.9 Barriers
2.4 CLASSICAL IPC PROBLEMS
2.4.1 The Dining Philosophers Problem
2.4.2 The Readers and Writers Problem
2.4.3 The Sleeping Barber Problem
2.5 SCHEDULING
2.5.1 Introduction to Scheduling
Process Behavior
When to Schedule
Categories of Scheduling Algorithms Scheduling Algorithm Goals
2.5.2 Scheduling in Batch Systems First-Come First-Served Shortest Job First
Shortest Remaining Time Next Three-Level Scheduling
2.5.3 Scheduling in Interactive Systems Round-Robin Scheduling
Priority Scheduling
Multiple Queues
Shortest Process Next Guaranteed Scheduling Lottery Scheduling Fair-Share Scheduling
2.5.4 Scheduling in Real-Time Systems
2.5.5 Policy versus Mechanism
2.5.6 Thread Scheduling
2.6 RESEARCH ON PROCESSES AND THREADS
2.7 SUMMARY
3 DEADLOCKS
3.1 RESOURCES
3.1.1 Preemptable and Nonpreemptable Resources
3.1.2 Resource Acquisition
3.2 INTRODUCTION TO DEADLOCKS
3.2.1 Conditions for Deadlock
3.2.2 Deadlock Modeling
3.3 THE OSTRICH ALGORITHM
3.4 DEADLOCK DETECTION AND RECOVERY
3.4.1 Deadlock Detection with One Resource of Each Type
3.4.2 Deadlock Detection with Multiple Resource of Each Type
3.4.3 Recovery from Deadlock
Recovery through Preemption Recovery through Rollback
Recovery through Killing Processes
3.5 DEADLOCK AVOIDANCE
3.5.1 Resource Trajectories
3.5.2 Safe and Unsafe States
3.5.3 The Banker’s Algorithm for a Single Resource
3.5.4 The Banker’s Algorithm for Multiple Resources
3.6 DEADLOCK PREVENTION
3.6.1 Attacking the Mutual Exclusion Condition
3.6.2 Attacking the Hold and Wait Condition
3.6.3 Attacking the No Preemption Condition
3.6.4 Attacking the Circular Wait Condition
3.7 OTHER ISSUES
3.7.1 Two-Phase Locking
3.7.2 Nonresource Deadlocks
3.7.3 Starvation
3.8 RESEARCH ON DEADLOCKS
3.9 SUMMARY
4 MEMORY MANAGEMENT
4.1 BASIC MEMORY MANAGEMENT
4.1.1 Monoprogramming without Swapping or Paging
4.1.2 Multiprogramming with Fixed Partitions
4.1.3 Modeling Multiprogramming
4.1.4 Analysis of Multiprogramming System Performance
4.1.5 Relocation and Protection
4.2 SWAPPING
4.2.1 Memory Management with Bitmaps
4.2.2 Memory Management with Linked Lists
4.3 VIRTUAL MEMORY
4.3.1 Paging
4.3.2 Page Tables Multilevel Page Tables
Structure of a Page Table Entry
4.3.3 TLBs—Translation Lookaside Buffers Software TLB Management
4.3.4 Inverted Page Tables
4.4 PAGE REPLACEMENT ALGORITHMS
4.4.1 The Optimal Page Replacement Algorithm
4.4.2 The Not Recently Used Page Replacement Algorithm
4.4.3 The First-In, First-Out (FIFO) Page Replacement Algorithm
4.4.4 The Second Chance Page Replacement Algorithm
4.4.5 The Clock Page Replacement Algorithm
4.4.6 The Least Recently Used (LRU) Page Replacement Algorithm
4.4.7 Simulating LRU in Software
4.4.8 The Working Set Page Replacement Algorithm
4.4.9 The WSClock Page Replacement Algorithm
4.4.10 Summary of Page Replacement Algorithms
4.5 MODELING PAGE REPLACEMENT ALGORITHMS
4.5.1 Belady’s Anomaly
4.5.2 Stack Algorithms
4.5.3 The Distance String
4.5.4 Predicting Page Fault Rates
4.6 DESIGN ISSUES FOR PAGING SYSTEMS
4.6.1 Local versus Global Allocation Policies
4.6.2 Load Control
4.6.3 Page Size
4.6.4 Separate Instruction and Data Spaces
4.6.5 Shared Pages
4.6.6 Cleaning Policy
4.6.7 Virtual Memory Interface
4.7 IMPLEMENTATION ISSUES
4.7.1 Operating System Involvement with Paging
4.7.2 Page Fault Handling
4.7.3 Instruction Backup
4.7.4 Locking Pages in Memory
4.7.5 Backing Store
4.7.6 Separation of Policy and Mechanism
4.8 SEGMENTATION
4.8.1 Implementation of Pure Segmentation
4.8.2 Segmentation with Paging: MULTICS
4.8.3 Segmentation with Paging: The Intel Pentium
4.9 RESEARCH ON MEMORY MANAGEMENT
4.10 SUMMARY
5 INPUT/OUTPUT
5.1 PRINCIPLES OF I/O HARDWARE
5.1.1 I/O Devices
5.1.2 Device Controllers
5.1.3 Memory-Mapped I/O
5.1.4 Direct Memory Access (DMA)
5.1.5 Interrupts Revisited
5.2 PRINCIPLES OF I/O SOFTWARE
5.2.1 Goals of the I/O Software
5.2.2 Programmed I/O
5.2.3 Interrupt-Driven I/O
5.2.4 I/O Using DMA
5.3 I/O SOFTWARE LAYERS
5.3.1 Interrupt Handlers
5.3.2 Device Drivers
5.3.3 Device-Independent I/O Software
Uniform Interfacing for Device Drivers Buffering
Error Reporting
Allocating and Releasing Dedicated Devices Device-Independent Block Size
5.3.4 User-Space I/O Software
5.4 DISKS
5.4.1 Disk Hardware Magnetic Disks
RAID CD-ROMs CD-Recordables CD-Rewritables DVD
5.4.2 Disk Formatting
5.4.3 Disk Arm Scheduling Algorithms
5.4.4 Error Handling
5.4.5 Stable Storage
5.5 CLOCKS
5.5.1 Clock Hardware
5.5.2 Clock Software
5.5.3 Soft Timers
5.6 CHARACTER-ORIENTED TERMINALS
5.6.1 RS-232 Terminal Hardware
5.6.2 Input Software
5.6.3 Output Software
5.7 GRAPHICAL USER INTERFACES
5.7.1 Personal Computer Keyboard, Mouse, and Display Hardware
5.7.2 Input Software
5.7.3 Output Software for Windows
BitMaps Fonts
5.8 NETWORK TERMINALS
5.8.1 The X Window System
5.8.2 The SLIM Network Terminal
5.9 POWER MANAGEMENT 5.9.1 Hardware Issues
5.9.2 Operating System Issues
The Display
The Hard Disk
The CPU
The Memory
Wireless Communication Thermal Management Battery Management Driver Interface
5.9.3 Degraded Operation
5.10 RESEARCH ON INPUT/OUTPUT
5.11 SUMMARY
6 FILE SYSTEMS
6.1 FILES
6.1.1 File Naming
6.1.2 File Structure
6.1.3 File Types
6.1.4 File Access
6.1.5 File Attributes
6.1.6 File Operations
6.1.7 An Example Program Using File System Calls
6.1.8 Memory-Mapped Files
6.2 DIRECTORIES
6.2.1 Single-Level Directory Systems
6.2.2 Two-level Directory Systems
6.2.3 Hierarchical Directory Systems
6.2.4 Path Names
6.2.5 Directory Operations
6.3 FILE SYSTEM IMPLEMENTATION
6.3.1 File System Layout
6.3.2 Implementing Files
Contiguous Allocation
Linked List Allocation
Linked List Allocation Using a Table in Memory I-nodes
6.3.3 Implementing Directories 6.3.4 Shared Files
6.3.5 Disk Space Management
Block Size
Keeping Track of Free Blocks Disk Quotas
6.3.6 File System Reliability Backups
File System Consistency
6.3.7 File System Performance Caching
Block Read Ahead Reducing Disk Arm Motion
6.3.8 Log-Structured File Systems
6.4 EXAMPLE FILE SYSTEMS
6.4.1 CD-ROM File Systems
The ISO 9660 File System Rock Ridge Extensions Joliet Extensions
6.4.2 The CP/M File System
6.4.3 The MS-DOS File System
6.4.4 The Windows 98 File System
6.4.5 The UNIX V7 File System
6.5 RESEARCH ON FILE SYSTEMS
6.6 SUMMARY
7 MULTIMEDIA OPERATING SYSTEMS
7.1 INTRODUCTION TO MULTIMEDIA
7.2 MULTIMEDIA FILES
7.2.1 Audio Encoding
7.2.2 Video Encoding
7.3 VIDEO COMPRESSION
7.3.1 The JPEG Standard
7.3.2 The MPEG Standard
7.4 MULTIMEDIA PROCESS SCHEDULING
7.4.1 Scheduling Homogeneous Processes
7.4.2 General Real-Time Scheduling
7.4.3 Rate Monotonic Scheduling
7.4.4 Earliest Deadline First Scheduling
7.5 MULTIMEDIA FILE SYSTEM PARADIGMS
7.5.1 VCR Control Functions
7.5.2 Near Video on Demand
7.5.3 Near Video on Demand with VCR Functions
7.6 FILE PLACEMENT
7.6.1 Placing a File on a Single Disk
7.6.2 Two Alternative File Organization Strategies
7.6.3 Placing Files for Near Video on Demand
7.6.4 Placing Multiple Files on a Single Disk
7.6.5 Placing Files on Multiple Disks
7.7 CACHING
7.7.1 Block Caching
7.7.2 File Caching
7.8 DISK SCHEDULING FOR MULTIMEDIA
7.8.1 Static Disk Scheduling
7.8.2 Dynamic Disk Scheduling
7.9 RESEARCH ON MULTIMEDIA
7.10 SUMMARY
8 MULTIPLE PROCESSOR SYSTEMS
8.1 MULTIPROCESSORS
8.1.1 Multiprocessor Hardware
UMA Bus-Based SMP Architectures
UMA Multiprocessors Using Crossbar Switches
UMA Multiprocessors Using Multistage Switching Networks NUMA Multiprocessors
8.1.2 Multiprocessor Operating System Types Each CPU Has Its Own Operating System Master-Slave Multiprocessors
Symmetric Multiprocessors
8.1.3 Multiprocessor Synchronization Spinning versus Switching
8.1.4 Multiprocessor Scheduling Timesharing
Space Sharing Gang Scheduling
8.2 MULTICOMPUTERS
8.2.1 Multicomputer Hardware
Interconnection Technology Network Interfaces
8.2.2 Low-Level Communication Software Node to Network Interface Communication
8.2.3 User-Level Communication Software Send and Receive
Blocking versus Nonblocking Calls
8.2.4 Remote Procedure Call Implementation Issues
8.2.5 Distributed Shared Memory Replication
False Sharing
Achieving Sequential Consistency
8.2.6 Multicomputer Scheduling
8.2.7 Load Balancing
A Graph-Theoretic Deterministic Algorithm
A Sender-Initiated Distributed Heuristic Algorithm
A Receiver-Initialed Distributed Heuristic Algorithm A Bidding Algorithm
8.3 DISTRIBUTED SYSTEMS
8.3.1 Network Hardware
Ethernet The Internet
8.3.2 Network Services and Protocols Network Services
Network Protocols
8.3.3 Document-Based Middleware
8.3.4 File System-Based Middleware
Transfer Model
The Directory Hierarchy Naming Transparency Semantics of File Sharing AFS
8.3.5 Shared Object-Based Middleware CORBA
Globe
8.3.6 Coordination-Based Middleware Linda
Publish/Subscribe Jini
8.4 RESEARCH ON MULTIPLE PROCESSOR SYSTEMS
8.5 SUMMARY
9 SECURITY
9.1 THE SECURITY ENVIRONMENT
9.1.1 Threats
9.1.2 Intruders
9.1.3 Accidental Data Loss
9.2 BASICS OF CRYPTOGRAPHY
9.2.1 Secret-Key Cryptography
9.2.2 Public-Key Cryptography
9.2.3 One-Way Functions
9.2.4 Digital Signatures
9.3 USER AUTHENTICATION
9.3.1 Authentication Using Passwords
How Crackers Break In
UNIX Password Security
Improving Password Security One-Time Passwords Challenge-Response Authentication
9.3.2 Authentication Using a Physical Object
9.3.3 Authentication Using Biometrics
9.3.4 Countermeasures
9.4 ATTACKS FROM INSIDE THE SYSTEM
9.4.1 Trojan Horses
9.4.2 Login Spoofing
9.4.3 Logic Bombs
9.4.4 Trap Doors
9.4.5 Buffer Overflow
9.4.6 Generic Security Attacks
9.4.7 Famous Security Flaws
Famous Security Flaws in UNIX Famous Security Flaws in TENEX Famous Security Flaws in OS/360
9.4.8 Design Principles for Security
9.5 ATTACKS FROM OUTSIDE THE SYSTEM
9.5.1 Virus Damage Scenarios
9.5.2 How Viruses Work
Companion Viruses Executable Program Viruses
Memory Resident Viruses Boot Sector Viruses Device Driver Viruses Macro Viruses
Source Code Viruses
9.5.3 How Viruses Spread
9.5.4 Antivirus and Anti-Antivirus Techniques
Virus Scanners
Integrity Checkers
Behavioral Checkers
Virus Avoidance
Recovery from a Virus Attack
9.5.5 The Internet Worm
9.5.6 Mobile Code
Sandboxing Interpretation Code signing
9.5.7 Java Security
9.6 PROTECTION MECHANISMS
9.6.1 Protection Domains
9.6.2 Access Control Lists
9.6.3 Capabilities
9.7 TRUSTED SYSTEMS
9.7.1 Trusted Computing Base
9.7.2 Formal Models of Secure Systems
9.7.3 Multilevel Security
The Bell-La Padula Model The Biba Model
9.7.4 Orange Book Security
9.7.5 Covert Channels
9.8 RESEARCH ON SECURITY
9.9 SUMMARY
10 CASE STUDY 1: UNIX AND LINUX 
10.1 HISTORY OF UNIX
10.1.1 UNICS
10.1.2 PDP-11 UNIX 
10.1.3 Portable UNIX 
10.1.4 Berkeley UNIX 
10.1.5 Standard UNIX 
10.1.6 MINIX
10.1.7 Linux
10.2 OVERVIEW OF UNIX 
10.2.1 UNIX Goals
10.2.2 Interfaces to UNIX 
10.2.3 The UNIX Shell
10.2.4 UNIX Utility Programs
10.2.5 Kernel Structure
10.3 PROCESSES IN UNIX
10.3.1 Fundamental Concepts
10.3.2 Process Management System Calls in UNIX
Thread Management System Calls
10.3.3 Implementation of Processes in UNIX Threads in UNIX
Threads in Linux Scheduling in UNIX Scheduling in Linux
10.3.4 Booting UNIX
10.4 MEMORY MANAGEMENT IN UNIX
10.4.1 Fundamental Concepts
10.4.2 Memory Management System Calls in UNIX
10.4.3 Implementation of Memory Management in UNIX
Swapping
Paging in UNIX
The Page Replacement Algorithm                               
Memory Management in Linux
10.5 INPUT/OUTPUT IN UNIX 10.5.1 Fundamental Concepts
Networking
10.5.2 Input/Output System Calls in UNIX
10.5.3 Implementation of Input/Output in UNIX
10.5.4 Streams
10.6 THE UNIX FILE SYSTEM
10.6.1 Fundamental Concepts
10.6.2 File System Calls in UNIX
10.6.3 Implementation of the UNIX File System
The Berkeley Fast File System The Linux File System
10.6.4 NFS: The Network File System NFS Architecture
NFS Protocols
NFS Implementation
10.7 SECURITY IN UNIX
10.7.1 Fundamental Concepts
10.7.2 Security System Calls in UNIX
10.7.3 Implementation of Security in UNIX
10.8 SUMMARY
11 CASE STUDY 2: WINDOWS 2000
11.1 HISTORY OF WINDOWS 2000
11.1.1 MS-DOS
11.1.2 Windows 95/98/Me
11.1.3 Windows NT
11.1.4 Windows 2000
11.2 PROGRAMMING WINDOWS 2000
11.2.1 The Win32 Application Programming Interface
11.2.2 The Registry
11.3 SYSTEM STRUCTURE
11.3.1 Operating System Structure
The Hardware Abstraction Layer The Kernel Layer
The Executive
The Device Drivers
11.3.2 Implementation of Objects The Object Name Space
11.3.3 Environment Subsystems
11.4 PROCESSES AND THREADS IN WINDOWS 2000
11.4.1 Fundamental Concepts
11.4.2 Job, Process, Thread and Fiber Management API Calls
Interprocess Communication
11.4.3 Implementation of Processes and Threads Scheduling
11.4.4 MS-DOS Emulation
11.4.5 Booting Windows 2000
11.5 MEMORY MANAGEMENT
11.5.1 Fundamental Concepts
11.5.2 Memory Management System Calls
11.5.3 Implementation of Memory Management
Page Fault Handling
The Page Replacement Algorithm Physical Memory Management
11.6 INPUT/OUTPUT IN WINDOWS 2000
11.6.1 Fundamental Concepts
11.6.2 Input/Output API Calls
11.6.3 Implementation of I/O                             
11.6.4 Device Drivers
11.7 THE WINDOWS 2000 FILE SYSTEM
11.7.1 Fundamental Concepts
11.7.2 File System API Calls in Windows 2000
11.7.3 Implementation of the Windows 2000 File System
File System Structure File Name Lookup File Compression File Encryption
11.8 SECURITY IN WINDOWS 2000
11.8.1 Fundamental Concepts
11.8.2 Security API Calls
11.8.3 Implementation of Security
11.9 CACHING IN WINDOWS 2000
11.10 SUMMARY
12 OPERATING SYSTEM DESIGN
12.1 THE NATURE OF THE DESIGN PROBLEM
12.1.1 Goals
12.1.2 Why is it Hard to Design an Operating System?
12.2 INTERFACE DESIGN
12.2.1 Guiding Principles
Simplicity Principle
Completeness Principle
Efficiency
12.2.2 Paradigms
User Interface Paradigms Execution Paradigms Data Paradigms
12.2.3 The System Call Interface
12.3 IMPLEMENTATION
12.3.1 System Structure Layered Systems
Exokernels Client-Server Systems Extensible Systems Kernel Threads
12.3.2 Mechanism versus Policy
12.3.3 Orthogonality
12.3.4 Naming
12.3.5 Binding Time
12.3.6 Static versus Dynamic Structures
12.3.7 Top-Down versus Bottom-Up Implementation
12.3.8 Useful Techniques
Hiding the Hardware Indirection Reusability Reentrancy
Brute Force
Check for Errors First
12.4 PERFORMANCE
12.4.1 Why Are Operating Systems Slow?
12.4.2 What Should Be Optimized?
12.4.3 Space-Time Trade-offs
12.4.4 Caching
12.4.5 Hints
12.4.6 Exploiting Locality
12.4.7 Optimize the Common Case
12.5 PROJECT MANAGEMENT
12.5.1 The Mythical Man Month
12.5.2 Team Structure
12.5.3 The Role of Experience
12.5.4 No Silver Bullet
12.6 TRENDS IN OPERATING SYSTEM DESIGN
12.6.1 Large Address Space Operating Systems
12.6.2 Networking
12.6.3 Parallel and Distributed Systems
12.6.4 Multimedia
12.6.5 Battery-Powered Computers
12.6.6 Embedded Systems
12.7 SUMMARY
13 READING LIST AND BIBLIOGRAPHY
13.1 SUGGESTIONS FOR FURTHER READING
13.1.1 Introduction and General Works
13.1.2 Processes and Threads
13.1.3 Deadlocks
13.1.4 Memory Management
13.1.5 Input/Output
13.1.6 File Systems
13.1.7 Multimedia Operating Systems
13.1.8 Multiple Processor Systems
13.1.9 Security
13.1.10 UNIX and Linux
13.1.11 Windows 2000
13.1.12 Design Principles
13.2 ALPHABETICAL BIBLIOGRAPHY
`

func getLevelOf(line string) int {
	if line == "" {
		return 0
	}
	startWithNumber := regexp.MustCompile(`^\d+`)
	if startWithNumber.MatchString(line) {
		return strings.Count(line, ".") + 1
	}
	return 4
}

func trimIndex(line string) string {
	numberOrDot := regexp.MustCompile(`^\d+`)
	if numberOrDot.MatchString(line) {
		i := 0
		for ; i < len(line); i++ {
			if line[i] == ' ' {
				break
			}
		}
		return line[i+1:]
	}
	return line
}

func parseToc() {
	tocRoot = &course.TreeNode{Name: "root", Path: ""}
	preNode := tocRoot
	for _, line := range strings.Split(toc, "\n") {
		//fmt.Print(preNode.Path, "\n")
		if line == "" {
			continue
		}
		level := getLevelOf(line)
		line = trimIndex(line)
		//fmt.Print("level: ", level, " line: ", line, "\n")
		for level != getLevelOf(preNode.Path)+1 {
			preNode = preNode.Parent
		}
		var childPath string
		if preNode.Path == "" {
			childPath = fmt.Sprint(len(preNode.Children) + 1)
		} else {
			childPath = preNode.Path + "." + fmt.Sprint(len(preNode.Children)+1)
		}
		child := &course.TreeNode{Name: line, Path: childPath, Parent: preNode}
		preNode.Children = append(preNode.Children, child)
		preNode = child
	}
}

func main() {
	courseMap = make(map[string]course.TreeNode)
	parseToc()
}
