import tkinter as tk
from tkinter import messagebox
import math
from collections import deque

class ScientificCalculator:
    def __init__(self, root):
        self.root = root
        self.root.title("Scientific Calculator")
        self.root.geometry("600x750")
        self.root.resizable(False, False)
        
        # Variables
        self.display_var = tk.StringVar(value="0")
        self.memory = 0
        self.history = deque(maxlen=20)
        self.last_result = 0
        
        # Setup UI
        self.setup_ui()
    
    def setup_ui(self):
        # Display
        display_frame = tk.Frame(self.root, bg="#2c3e50")
        display_frame.pack(fill=tk.BOTH, padx=10, pady=10)
        
        self.display = tk.Entry(
            display_frame, 
            textvariable=self.display_var,
            font=("Arial", 28, "bold"),
            justify="right",
            bg="#ecf0f1",
            fg="#2c3e50",
            bd=0
        )
        self.display.pack(fill=tk.BOTH, ipady=20)
        
        # Memory indicator
        self.memory_label = tk.Label(
            self.root,
            text="",
            font=("Arial", 10),
            fg="#e74c3c"
        )
        self.memory_label.pack(anchor="w", padx=15)
        
        # Buttons frame
        buttons_frame = tk.Frame(self.root, bg="#34495e")
        buttons_frame.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)
        
        # Button layout
        buttons = [
            # Row 1 - Memory
            [("MC", self.memory_clear, "#e74c3c"), ("MR", self.memory_recall, "#e74c3c"), 
             ("M+", self.memory_add, "#e74c3c"), ("M-", self.memory_subtract, "#e74c3c")],
            
            # Row 2 - Scientific
            [("sin", self.sin_func, "#9b59b6"), ("cos", self.cos_func, "#9b59b6"), 
             ("tan", self.tan_func, "#9b59b6"), ("π", self.insert_pi, "#9b59b6")],
            
            # Row 3 - Scientific
            [("log", self.log_func, "#9b59b6"), ("ln", self.ln_func, "#9b59b6"), 
             ("√", self.sqrt_func, "#9b59b6"), ("!", self.factorial, "#9b59b6")],
            
            # Row 4 - Powers
            [("x²", self.square, "#16a085"), ("x³", self.cube, "#16a085"), 
             ("xʸ", self.power, "#16a085"), ("1/x", self.reciprocal, "#16a085")],
            
            # Row 5 - Clear and basic
            [("C", self.clear, "#e67e22"), ("CE", self.clear_entry, "#e67e22"), 
             ("←", self.backspace, "#e67e22"), ("÷", self.insert_operator, "#e67e22")],
            
            # Row 6 - Numbers and operators
            [("7", self.insert_number, "#34495e"), ("8", self.insert_number, "#34495e"), 
             ("9", self.insert_number, "#34495e"), ("×", self.insert_operator, "#e67e22")],
            
            # Row 7
            [("4", self.insert_number, "#34495e"), ("5", self.insert_number, "#34495e"), 
             ("6", self.insert_number, "#34495e"), ("-", self.insert_operator, "#e67e22")],
            
            # Row 8
            [("1", self.insert_number, "#34495e"), ("2", self.insert_number, "#34495e"), 
             ("3", self.insert_number, "#34495e"), ("+", self.insert_operator, "#e67e22")],
            
            # Row 9
            [("0", self.insert_number, "#34495e"), (".", self.insert_number, "#34495e"), 
             ("(", self.insert_operator, "#e67e22"), (")", self.insert_operator, "#e67e22")],
            
            # Row 10
            [("=", self.calculate, "#27ae60")],
        ]
        
        for row in buttons:
            row_frame = tk.Frame(buttons_frame, bg="#34495e")
            row_frame.pack(fill=tk.BOTH, expand=True, pady=5)
            
            for button_info in row:
                text, command, color = button_info
                btn = tk.Button(
                    row_frame,
                    text=text,
                    command=lambda cmd=command, t=text: cmd(t),
                    font=("Arial", 14, "bold"),
                    bg=color,
                    fg="white",
                    border=0,
                    activebackground="#2c3e50",
                    activeforeground="white",
                    padx=20,
                    pady=15
                )
                btn.pack(side=tk.LEFT, fill=tk.BOTH, expand=True, padx=2)
    
    # Number and operator insertion
    def insert_number(self, char):
        current = self.display_var.get()
        if current == "0" and char != ".":
            self.display_var.set(char)
        elif char == "." and "." in current:
            return
        else:
            self.display_var.set(current + char)
    
    def insert_operator(self, op):
        current = self.display_var.get()
        if op in "()":
            self.display_var.set(current + op)
        else:
            self.display_var.set(current + " " + op + " ")
    
    def insert_pi(self, _):
        self.display_var.set(self.display_var.get() + str(math.pi))
    
    # Scientific functions
    def sin_func(self, _):
        try:
            val = math.radians(float(self.display_var.get()))
            self.display_var.set(str(math.sin(val)))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def cos_func(self, _):
        try:
            val = math.radians(float(self.display_var.get()))
            self.display_var.set(str(math.cos(val)))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def tan_func(self, _):
        try:
            val = math.radians(float(self.display_var.get()))
            self.display_var.set(str(math.tan(val)))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def log_func(self, _):
        try:
            self.display_var.set(str(math.log10(float(self.display_var.get()))))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def ln_func(self, _):
        try:
            self.display_var.set(str(math.log(float(self.display_var.get()))))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def sqrt_func(self, _):
        try:
            self.display_var.set(str(math.sqrt(float(self.display_var.get()))))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def square(self, _):
        try:
            val = float(self.display_var.get())
            self.display_var.set(str(val ** 2))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def cube(self, _):
        try:
            val = float(self.display_var.get())
            self.display_var.set(str(val ** 3))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def power(self, _):
        self.display_var.set(self.display_var.get() + " ** ")
    
    def reciprocal(self, _):
        try:
            val = float(self.display_var.get())
            if val == 0:
                messagebox.showerror("Error", "Cannot divide by zero")
            else:
                self.display_var.set(str(1 / val))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def factorial(self, _):
        try:
            val = int(float(self.display_var.get()))
            if val < 0:
                messagebox.showerror("Error", "Factorial of negative number")
            else:
                self.display_var.set(str(math.factorial(val)))
        except:
            messagebox.showerror("Error", "Invalid input")
    
    # Calculation
    def calculate(self, _):
        try:
            expression = self.display_var.get()
            result = eval(expression)
            self.history.append(f"{expression} = {result}")
            self.last_result = result
            self.display_var.set(str(result))
        except:
            messagebox.showerror("Error", "Invalid expression")
    
    # Clearing
    def clear(self, _):
        self.display_var.set("0")
        self.memory = 0
        self.update_memory_label()
    
    def clear_entry(self, _):
        self.display_var.set("0")
    
    def backspace(self, _):
        current = self.display_var.get()
        self.display_var.set(current[:-1] if len(current) > 1 else "0")
    
    # Memory functions
    def memory_add(self, _):
        try:
            self.memory += float(self.display_var.get())
            self.update_memory_label()
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def memory_subtract(self, _):
        try:
            self.memory -= float(self.display_var.get())
            self.update_memory_label()
        except:
            messagebox.showerror("Error", "Invalid input")
    
    def memory_recall(self, _):
        self.display_var.set(str(self.memory))
    
    def memory_clear(self, _):
        self.memory = 0
        self.update_memory_label()
    
    def update_memory_label(self):
        self.memory_label.config(text=f"M: {self.memory}" if self.memory != 0 else "")

if __name__ == "__main__":
    root = tk.Tk()
    calc = ScientificCalculator(root)
    root.mainloop()
